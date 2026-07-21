import os
import glob
import xgboost as xgb
from flask import Flask, request, jsonify
from flask_cors import CORS
import re
import time
import threading
import traceback
from openai import OpenAI

# 设置项目根目录
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
MODEL_DIR = os.path.join(BASE_DIR, "crop_models")
FEATURE_NAMES = ["N", "P", "K", "temperature", "humidity", "ph", "rainfall", "pesticide"]
print(f"📂 模型目录路径: {MODEL_DIR}")

# 验证模型目录
if not os.path.exists(MODEL_DIR):
    print(f"❌ 错误: 模型目录不存在: {MODEL_DIR}")
    exit(1)
else:
    print(f"✅ 模型目录存在")
    print(f"🔍 目录内容:")
    for item in os.listdir(MODEL_DIR)[:5]:
        print(f"  - {item}")
    if len(os.listdir(MODEL_DIR)) > 5:
        print(f"  显示前5项，共 {len(os.listdir(MODEL_DIR))} 个文件")

# 初始化大模型客户端
AI_CLIENT = OpenAI(
    api_key="your-api-key-here",  # 替换为你的 API Key
    base_url="https://dashscope.aliyuncs.com/compatible-mode/v1"
)
AI_MODEL = "deepseek-r1"  # 使用DeepSeek模型

app = Flask(__name__)
# 允许所有来源的跨域请求
CORS(app, resources={r"/*": {"origins": "*"}})


# 解决跨域问题的方法
@app.after_request
def after_request(response):
    # 允许所有来源的跨域请求
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
    response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS')
    return response


# 中英文作物名映射（与前端完全一致）
CROP_MAPPING = {
    "苹果": "Apples",
    "香蕉": "Bananas",
    "大麦": "Barley",
    "蓝莓": "Blueberries",
    "荞麦": "Buckwheat",
    "卷心菜": "Cabbages",
    "胡萝卜和萝卜": "Carrots_and_turnips",
    "樱桃": "Cherries",
    "干辣椒": "Chillies_and_peppers_dry",
    "黄瓜和小黄瓜": "Cucumbers_and_gherkins",
    "茄子": "Eggplants",
    "生姜（鲜）": "Ginger_raw",
    "葡萄": "Grapes",
    "青玉米": "Green_corn",
    "青蒜": "Green_garlic",
    "花生": "Groundnuts",
    "猕猴桃": "Kiwi_fruit",
    "韭菜和其他葱属蔬菜": "Leeks_and_other_alliaceous_vegetables",
    "生菜和菊苣": "Lettuce_and_chicory",
    "玉米": "Maize",
    "芒果、番石榴和山竹": "Mangoes_guavas_and_mangosteens",
    "小米": "Millet",
    "燕麦": "Oats",
    "洋葱和青葱（干的）": "Onions_and_shallots_dry",
    "橙子": "Oranges",
    "梨": "Pears",
    "青豆": "Peas_green",
    "柿子": "Persimmons",
    "菠萝": "Pineapples",
    "柚子和葡萄柚": "Pomelos_and_grapefruits",
    "马铃薯": "Potatoes",
    "南瓜、西葫芦和葫芦": "Pumpkins_squash_and_gourds",
    "芝麻": "Sesame_seed",
    "高粱": "Sorghum",
    "大豆": "Soya_beans",
    "菠菜": "Spinach",
    "草莓": "Strawberries",
    "甜菜": "Sugar_beet",
    "甘蔗": "Sugar_cane",
    "甘薯": "Sweet_potatoes",
    "番茄": "Tomatoes",
    "黑小麦": "Triticale",
    "西瓜": "Watermelons",
}

# 模型名称规范化和反向映射
MODEL_NAME_MAP = {}
REVERSE_MODEL_MAP = {}
for chinese_name, english_name in CROP_MAPPING.items():
    # 标准化名称（去除空格和特殊字符）
    base_name = re.sub(r'[^a-zA-Z0-9]', '_', english_name)

    # 建立映射
    MODEL_NAME_MAP[chinese_name] = base_name
    MODEL_NAME_MAP[english_name] = base_name
    MODEL_NAME_MAP[english_name.replace('_', ' ')] = base_name

    # 建立反向映射
    REVERSE_MODEL_MAP[base_name] = chinese_name


# 按需加载模型的函数
def load_specific_model(crop_name):
    """动态加载指定作物的模型"""
    try:
        print(f"🔍 尝试加载模型: {crop_name}")

        # 转换名称格式
        if crop_name in MODEL_NAME_MAP:
            base_name = MODEL_NAME_MAP[crop_name]
        else:
            # 去除空格和特殊字符
            base_name = re.sub(r'[^a-zA-Z0-9]', '_', crop_name)
            if base_name not in REVERSE_MODEL_MAP:
                # 尝试匹配模型文件
                all_files = glob.glob(os.path.join(MODEL_DIR, "*.json"))
                for file_path in all_files:
                    file_base = os.path.basename(file_path)
                    model_name = re.sub(r'_model\.json$', '', file_base)
                    if crop_name.lower() in model_name.lower():
                        base_name = model_name
                        print(f"🔀 匹配到模型文件: {model_name} -> {file_path}")
                        break

        model_path = os.path.join(MODEL_DIR, f"{base_name}_model.json")

        # 验证文件是否存在
        if not os.path.exists(model_path):
            # 尝试匹配模型文件
            matching_files = glob.glob(os.path.join(MODEL_DIR, f"*{base_name}*.json"))
            if matching_files:
                model_path = matching_files[0]
                print(f"🔀 使用替代文件: {model_path}")
            else:
                print(f"❌ 模型文件不存在: {base_name}_model.json")
                return None

        start_time = time.time()
        model = xgb.Booster()
        model.load_model(model_path)
        load_time = time.time() - start_time

        print(f"✅ 成功加载模型: {base_name} ({load_time:.3f}s)")
        return model
    except Exception as e:
        print(f"❌ 加载模型失败: {crop_name} - {str(e)}")
        traceback.print_exc()
        return None


def generate_ai_feedback(crop_name, yield_value, features):
    """使用大模型生成智能反馈和建议"""
    try:
        # 准备提示词
        prompt = f"""
你是一位专业的农业专家，请根据以下信息为农民提供种植建议：
- 作物类型: {crop_name}
- 预测产量: {yield_value:.2f} 公斤/公顷
- 土壤氮含量(N): {features.get('N', 0)} kg/亩
- 土壤磷含量(P): {features.get('P', 0)} kg/亩
- 土壤钾含量(K): {features.get('K', 0)} kg/亩
- 年均温度: {features.get('temperature', 0)}℃
- 湿度: {features.get('humidity', 0)}%
- 土壤pH值: {features.get('ph', 0)}
- 年降雨量: {features.get('rainfall', 0)}mm
- 农药用量: {features.get('pesticide', 0)}kg/ha

请完成以下任务：
1. 根据预测产量给出评级（优秀/良好/一般/较差/很差）
2. 提供1-2句产量评价，至少20字
3. 提供3条具体的种植建议，每条建议至少50-60字
4. 建议应基于输入参数，具有针对性

请使用以下JSON格式返回结果：
{{
  "rating": "评级",
  "comment": "产量评价",
  "suggestions": ["建议1", "建议2", "建议3"]
}}
"""

        # 调用大模型API - 使用qwen-plus模型
        response = AI_CLIENT.chat.completions.create(
            model="qwen-plus",  # 使用qwen-plus模型
            messages=[
                {"role": "system", "content": "你是一位专业的农业专家，为农民提供种植建议。"},
                {"role": "user", "content": prompt}
            ],
            temperature=0.3,
            max_tokens=500
        )

        # 解析响应
        content = response.choices[0].message.content

        # 尝试提取JSON部分
        try:
            # 查找JSON格式的内容
            start_idx = content.find('{')
            end_idx = content.rfind('}') + 1
            json_str = content[start_idx:end_idx]

            # 解析JSON
            import json
            result = json.loads(json_str)

            # 验证结构
            if "rating" in result and "comment" in result and "suggestions" in result:
                return result
            else:
                print("⚠️ 大模型返回格式不完整，使用默认建议")
        except Exception as e:
            print(f"❌ 解析大模型响应失败: {str(e)}")
            print(f"原始响应: {content}")

        # 如果解析失败，使用默认建议
        return generate_default_feedback(yield_value, crop_name)
    except Exception as e:
        print(f"❌ 调用大模型失败: {str(e)}")
        traceback.print_exc()
        return generate_default_feedback(yield_value, crop_name)


def generate_default_feedback(yield_value, crop_name):
    """生成默认反馈信息"""
    if yield_value > 8000:
        rating = "优秀"
        comment = f"{crop_name}产量非常高，种植条件非常理想"
    elif yield_value > 6000:
        rating = "良好"
        comment = f"{crop_name}产量良好，有提升空间"
    elif yield_value > 4000:
        rating = "一般"
        comment = f"{crop_name}产量一般，需要改进种植方案"
    elif yield_value > 2000:
        rating = "较差"
        comment = f"{crop_name}产量较低，种植条件不理想"
    else:
        rating = "很差"
        comment = f"{crop_name}产量非常低，建议更换种植方案"

    # 生成建议
    suggestions = [
        f"保持合理的{crop_name}种植密度",
        f"定期监测{crop_name}病虫害情况",
        f"根据生长阶段调整{crop_name}水肥管理"
    ]
    if rating in ["一般", "较差", "很差"]:
        suggestions.insert(0, f"增加{crop_name}的肥料投入")
        suggestions.insert(0, f"优化{crop_name}的灌溉系统")

    return {
        "rating": rating,
        "comment": comment,
        "suggestions": suggestions[:3]  # 最多3条建议
    }


class CropPredictor:
    def __init__(self):
        self.models_cache = {}  # 缓存已加载的模型

    def predict(self, crop_name, features):
        """按需加载模型并进行预测"""
        print(f"🌱 开始预测: {crop_name}")

        # 尝试从缓存获取模型
        if crop_name in self.models_cache:
            print(f"♻️ 使用缓存模型: {crop_name}")
            return self.do_prediction(crop_name, features)

        # 尝试标准化名称后获取
        normalized_name = self.normalize_crop_name(crop_name)
        if normalized_name in self.models_cache:
            print(f"🔀 使用标准化名称: {normalized_name}")
            return self.do_prediction(normalized_name, features)

        # 加载模型
        model = load_specific_model(crop_name)
        if model is None:
            # 尝试使用标准化名称加载
            normalized_name = self.normalize_crop_name(crop_name)
            model = load_specific_model(normalized_name)

            if model is None:
                # 生成可用作物列表
                all_files = glob.glob(os.path.join(MODEL_DIR, "*.json"))
                available = [re.sub(r'_model\.json$', '', os.path.basename(f)) for f in all_files]

                # 尝试找到相似名称
                from difflib import get_close_matches
                similar = get_close_matches(crop_name, available, n=3, cutoff=0.6)

                return {
                    "error": f"未找到作物模型: '{crop_name}'",
                    "suggestions": [f"请尝试: {s}" for s in similar] if similar else [],
                    "available_crops": available
                }

        # 添加到缓存
        self.models_cache[crop_name] = model
        return self.do_prediction(crop_name, features)

    def normalize_crop_name(self, name):
        """标准化作物名称"""
        # 去除空格
        normalized = name.replace(" ", "_")
        # 规范化下划线
        normalized = re.sub(r'_+', '_', normalized)
        # 统一大小写
        normalized = normalized.lower()
        return normalized

    def do_prediction(self, crop_name, features):
        """执行预测"""
        model = self.models_cache.get(crop_name)
        if model is None:
            return {
                "error": f"模型加载异常: '{crop_name}'",
                "available_crops": list(self.models_cache.keys())
            }

        # 准备特征数据 (确保顺序正确)
        input_data = []
        for col in FEATURE_NAMES:
            value = features.get(col, 0)
            try:
                input_data.append(float(value))
            except (TypeError, ValueError):
                input_data.append(0.0)
                print(f"⚠️ 转换特征 '{col}' 失败，使用默认值0.0")

        # 创建DMatrix并设置特征名称
        try:
            dmatrix = xgb.DMatrix([input_data], feature_names=FEATURE_NAMES)
            yield_value = float(model.predict(dmatrix)[0])

            print(f"📊 预测成功: {crop_name} -> {yield_value:.2f} kg/ha")

            # 使用大模型生成反馈和建议
            feedback = generate_ai_feedback(crop_name, yield_value, features)

            return {
                "crop": crop_name,
                "yield": round(yield_value, 2),
                "rating": feedback["rating"],
                "comment": feedback["comment"],
                "suggestions": feedback["suggestions"]
            }
        except Exception as e:
            print(f"❌ 预测失败: {crop_name} - {str(e)}")
            traceback.print_exc()
            return {
                "error": f"预测过程中出错: {str(e)}",
                "crop": crop_name,
                "features": features
            }


# 初始化预测器
print("\n" + "=" * 60)
print("🚀 启动作物产量预测服务 (智能建议版)")
print("=" * 60)
predictor = CropPredictor()


@app.route('/predict', methods=['POST'])
def predict_endpoint():
    """预测接口 - 使用大模型生成建议"""
    try:
        data = request.get_json()
        if not data:
            return jsonify({
                "error": "请求数据为空"
            }), 400
        if "crop" not in data:
            return jsonify({
                "error": "缺少参数: crop"
            }), 400
        if "features" not in data:
            return jsonify({
                "error": "缺少参数: features"
            }), 400

        # 前端可能发送中文或英文作物名
        crop_name = data["crop"]
        features = data["features"]

        # 记录更详细的请求信息
        print(f"🌱 预测请求: [{crop_name}]")
        # 完整记录所有特征值
        feature_str = ", ".join([
            f"{key}={features.get(key, 0)}"
            for key in FEATURE_NAMES
        ])
        print(f"  特征: {feature_str}")

        result = predictor.predict(crop_name, features)
        return jsonify(result)
    except Exception as e:
        traceback.print_exc()
        return jsonify({
            "error": f"服务器错误: {str(e)}"
        }), 500


@app.route('/health', methods=['GET'])
def health_check():
    """健康检查接口"""
    return jsonify({
        "status": "ready",
        "models_loaded": len(predictor.models_cache),
        "supported_features": FEATURE_NAMES,
        "ai_model": AI_MODEL
    })


if __name__ == '__main__':
    print("\n" + "=" * 60)
    print(f"🌐 服务运行在: http://localhost:5000")
    print(f"   测试接口: http://localhost:5000/health")
    print(f"   使用大模型: {AI_MODEL}")
    print("=" * 60)

    # 确保必要的模型特征名称正确
    print(f"🔠 必需特征名称: {FEATURE_NAMES}")

    # 预加载几个常见模型
    for crop in ["Apples", "Grapes", "小麦", "玉米"]:
        load_specific_model(crop)

    app.run(host='0.0.0.0', port=8901, debug=False, threaded=True)