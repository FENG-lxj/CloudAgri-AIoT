from flask import Flask, request, jsonify
from flask_cors import CORS
from http import HTTPStatus
from dashscope import Application

# ===== 百炼 API 配置 =====
API_KEY = "your-api-key-here"
APP_ID = "your-app-id-here"

# ===== Flask 应用 =====
app = Flask(__name__)
CORS(app)  # 解决跨域问题


@app.route('/IoTA/ai/AgricultureExpertAI', methods=['GET'])
def agriculture_expert_ai():
    user_query = request.args.get("Issue", "").strip()
    if not user_query:
        return jsonify({
            "code": 400,
            "message": "缺少问题内容",
            "data": {"AIReply": ""}
        })

    response = Application.call(
        api_key=API_KEY,
        app_id=APP_ID,
        prompt=user_query,
        stream=False  # 关闭流式调用，等待完整结果返回
    )

    if response.status_code == HTTPStatus.OK:
        ai_reply = response.output.text
        return jsonify({
            "code": 200,
            "message": "成功",
            "data": {"AIReply": ai_reply}
        })
    else:
        return jsonify({
            "code": response.status_code,
            "message": response.message,
            "data": {"AIReply": ""}
        })


# ===== 启动服务 =====
if __name__ == "__main__":
    app.run(host="0.0.0.0", port=4999, debug=True)
