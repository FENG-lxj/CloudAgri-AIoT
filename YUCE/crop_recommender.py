import os
import glob
import numpy as np
import pandas as pd
from flask import Flask, request, jsonify
from flask_cors import CORS
import traceback

# 设置项目根目录
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
DATA_DIR = os.path.join(BASE_DIR, "crop_data_china")  # 修改为数据目录
FEATURE_NAMES = ["N", "P", "K", "temperature", "humidity", "ph", "rainfall"]

# 验证数据目录
if not os.path.exists(DATA_DIR):
    print(f"❌ 错误: 数据目录不存在: {DATA_DIR}")
    exit(1)

app = Flask(__name__)
CORS(app, resources={r"/*": {"origins": "*"}})


# 解决跨域问题
@app.after_request
def after_request(response):
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
    response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS')
    return response


# 作物映射表（39种作物）
CROP_MAPPING = {
    "苹果": "Apples",
    "香蕉": "Bananas",
    "大麦": "Barley",
    "蓝莓": "Blueberries",
    "荞麦": "Buckwheat",
    "卷心菜": "Cabbages",
    "胡萝卜和萝卜": "Carrots_and_turnips",
    "樱桃": "Cherries",
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
    "小米": "Millet",
    "燕麦": "Oats",
    "橙子": "Oranges",
    "梨": "Pears",
    "青豆": "Peas_green",
    "柿子": "Persimmons",
    "菠萝": "Pineapples",
    "柚子和葡萄柚": "Pomelos_and_grapefruits",
    "马铃薯": "Potatoes",
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
    "小麦": "Wheat"
}


class CropRecommender:
    def __init__(self):
        self.crop_data = {}
        self.load_all_crop_data()

    def load_all_crop_data(self):
        """加载所有作物的特征数据"""
        print("🌱 开始加载所有作物特征数据...")

        # 获取所有数据文件
        data_files = glob.glob(os.path.join(DATA_DIR, "*.csv"))
        if not data_files:
            print("❌ 错误: 没有找到任何数据文件")
            return

        # 加载所有作物数据
        for file_path in data_files:
            try:
                # 从文件名提取作物名称
                file_name = os.path.basename(file_path)
                crop_english = file_name.replace(".csv", "")

                # 查找对应的中文名称
                chinese_name = None
                for cn, en in CROP_MAPPING.items():
                    if en == crop_english:
                        chinese_name = cn
                        break

                if not chinese_name:
                    print(f"⚠️ 未找到映射: {crop_english}")
                    continue

                # 加载数据并计算平均特征值
                df = pd.read_csv(file_path)
                avg_features = {
                    "N": df["N"].mean(),
                    "P": df["P"].mean(),
                    "K": df["K"].mean(),
                    "temperature": df["temperature"].mean(),
                    "humidity": df["humidity"].mean(),
                    "ph": df["ph"].mean(),
                    "rainfall": df["rainfall"].mean()
                }

                # 添加到作物数据字典
                self.crop_data[chinese_name] = {
                    "english_name": crop_english,
                    "features": avg_features
                }

                print(f"✅ 已加载: {chinese_name}")
            except Exception as e:
                print(f"❌ 加载失败: {file_name} - {str(e)}")

        print(f"📊 数据加载完成: 成功 {len(self.crop_data)} 种作物")

    def recommend_crop(self, input_features):
        """根据输入特征推荐最相似的作物"""
        if not self.crop_data:
            return {"error": "没有可用的作物数据"}

        # 验证输入特征
        for feature in FEATURE_NAMES:
            if feature not in input_features:
                return {"error": f"缺少必要特征: {feature}"}

        # 计算与每种作物的相似度（欧氏距离）
        similarities = {}
        for crop_name, crop_info in self.crop_data.items():
            distance = 0
            for feature in FEATURE_NAMES:
                # 标准化特征值（避免量纲影响）
                input_val = float(input_features[feature])
                crop_val = crop_info["features"][feature]

                # 计算归一化后的差值平方
                if feature in ["N", "P", "K"]:
                    norm_diff = (input_val - crop_val) / 100
                elif feature == "temperature":
                    norm_diff = (input_val - crop_val) / 50
                elif feature == "humidity":
                    norm_diff = (input_val - crop_val) / 100
                elif feature == "ph":
                    norm_diff = (input_val - crop_val) / 14
                elif feature == "rainfall":
                    norm_diff = (input_val - crop_val) / 2000

                distance += norm_diff ** 2

            # 存储相似度（距离越小越相似）
            similarities[crop_name] = np.sqrt(distance)

        # 按相似度排序（从最相似到最不相似）
        sorted_crops = sorted(similarities.items(), key=lambda x: x[1])

        # 获取推荐作物
        recommended = sorted_crops[0]
        top_crops = sorted_crops[1:6]  # 前5个备选作物

        return {
            "recommended_crop": recommended[0],
            "recommended_crop_english": self.crop_data[recommended[0]]["english_name"],
            "similarity_score": 1 / (1 + recommended[1]),  # 转换为相似度分数（0-1）
            "top_crops": [
                {
                    "name": crop[0],
                    "english_name": self.crop_data[crop[0]]["english_name"],
                    "similarity_score": 1 / (1 + crop[1])
                } for crop in top_crops
            ],
            "input_features": input_features  # 返回输入的特征用于前端显示
        }


# 初始化推荐器
print("\n" + "=" * 60)
print("🚀 启动作物推荐服务")
print("=" * 60)
recommender = CropRecommender()


@app.route('/recommend_crop', methods=['POST'])
def recommend_crop():
    """推荐最适合的作物"""
    try:
        data = request.get_json()
        if not data or "features" not in data:
            return jsonify({"error": "无效请求"}), 400

        features = data["features"]

        # 记录完整的请求参数
        print(f"🌱 推荐请求: N={features.get('N')}, P={features.get('P')}, K={features.get('K')}, "
              f"temperature={features.get('temperature')}, humidity={features.get('humidity')}, "
              f"ph={features.get('ph')}, rainfall={features.get('rainfall')}")

        result = recommender.recommend_crop(features)
        return jsonify(result)
    except Exception as e:
        traceback.print_exc()
        return jsonify({"error": f"服务器错误: {str(e)}"}), 500


@app.route('/health', methods=['GET'])
def health_check():
    """健康检查接口"""
    return jsonify({
        "status": "ready",
        "crops_loaded": len(recommender.crop_data),
        "supported_features": FEATURE_NAMES
    })


@app.route('/available_crops', methods=['GET'])
def available_crops():
    """获取可用作物列表"""
    return jsonify({
        "crops": list(recommender.crop_data.keys())
    })


if __name__ == '__main__':
    print("\n" + "=" * 60)
    print(f"🌐 服务运行在: http://localhost:5001")
    print(f"   测试接口: http://localhost:5001/health")
    print("=" * 60)
    app.run(host='localhost', port=8900, debug=False, threaded=True)