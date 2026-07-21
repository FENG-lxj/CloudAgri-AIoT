import json

import cv2
import subprocess
import threading
import time
from datetime import datetime, timedelta
from collections import defaultdict
from ultralytics import YOLO
import requests
import random
import string
import paho.mqtt.client as mqtt

# ====================== 用户可配置参数 ======================
# RTMP_URL = "rtmp://39.102.208.105:1935/PastoralCloud/monitor_1"  # 推流地址
RTMP_URL = "rtmp://127.0.0.1/live/livestream"
# MODEL_PATH = "runs/detect/train/weights/best.pt"  # 训练好的模型路径
MODEL_PATH = "runs/detect/train/weights/best.pt"  # 训练好的模型路径（请替换为实际路径）
TARGET_FPS = 5  # 目标检测帧率（手动调整）
CAMERA_INDEX = 0  # 摄像头设备索引（0为默认摄像头）
CAMERA_RESOLUTION = (640, 480)  # 摄像头分辨率（需与FFmpeg设置一致）

ICID = 1  # 摄像头设备ID
TOKEN = "Bearer your-jwt-token-here"  # 认证Token，从系统登录接口获取
API_URL = "http://127.0.0.1:20201/IoTA/users/AiImageInfo"  # API地址
# API_URL = "http://39.102.208.105:20201/IoTA/users/AiImageInfo"  # API地址
CLASS_HIERARCHY = {  # 分类映射字典（一级分类、二级分类、三级分类）
    "person": (2, "人物入侵", None),
    "fire": (3, None, None),
    "Spodoptera_frugiperda": (1, "玉米", "草地贪夜蛾"),
    "Blast": (1, "水稻", "稻瘟病"),
    "Blight": (1, "水稻", "疫病"),
    "Brownspot": (1, "水稻", "褐斑病"),
}

topic = "Go_IoT_Logs"
ICHB_topic = "Go_IoT_ICHB"
broker = "127.0.0.1"
port = 1883
keepalive = 60  # 保持连接的时间（秒）


# ===========================================================


def init_ffmpeg_stream():
    """初始化FFmpeg推流进程（固定输出帧率为TARGET_FPS）"""
    ffmpeg_cmd = [
        '/opt/homebrew/bin/ffmpeg',
        '-y', '-f', 'rawvideo', '-vcodec', 'rawvideo',
        '-pix_fmt', 'bgr24',
        '-s', f'{CAMERA_RESOLUTION[0]}x{CAMERA_RESOLUTION[1]}',  # 分辨率
        '-r', str(TARGET_FPS),  # 输出帧率=目标检测帧率
        '-i', '-',  # 从标准输入读取
        '-c:v', 'libx264',
        '-preset', 'ultrafast',  # 编码速度优先
        '-tune', 'zerolatency',  # 降低延迟
        '-f', 'flv',
        RTMP_URL
    ]
    return subprocess.Popen(ffmpeg_cmd, stdin=subprocess.PIPE)


def process_frame(frame, model):
    """执行目标检测并返回标注后的帧和检测结果"""
    results = model.predict(frame, conf=0.5, verbose=False)  # verbose关闭日志
    annotated_frame = results[0].plot()  # 自动绘制检测框
    return annotated_frame, results[0].boxes


def upload_image(image, class_name):
    """多线程上传图片到服务器（5分钟内同类别不上传重复图片）"""

    def _upload():
        try:
            # 获取分类层级信息
            if class_name not in CLASS_HIERARCHY:
                print(f"未配置的分类: {class_name}")
                return

            level1, level2, level3 = CLASS_HIERARCHY[class_name]

            # 生成识别ID（示例规则实现）
            now = datetime.now()
            random_char = random.choice(
                string.ascii_uppercase)  ####################################################################################################################################################################

            # 生成识别前缀ID
            if len(class_name) > 5:
                class_id = class_name[:5].upper()  # 如果超过5位，取前5位并转大写
            else:
                class_id = class_name.upper()  # 如果不足5位，全部转大写

            identify_id = (
                f"{class_id}{ICID}"
                f"{now.strftime('%Y%m%d%H%M%S')}"
                f"{random_char}"
            )

            # 转换图片为JPEG
            _, buffer = cv2.imencode('.jpg', image)
            jpg_bytes = buffer.tobytes()
            files = {'AiImage': ('detection.jpg', jpg_bytes, 'image/jpeg')}

            # 构造请求参数
            data = {
                "ICID": ICID,
                "Level1Class": level1,
                "IdentifyID": identify_id,
                "Timestamp": int(time.time())
            }
            if level2: data["Level2Class"] = level2
            if level3: data["Level3Class"] = level3

            # 发送请求
            headers = {"Authorization": TOKEN}
            response = requests.post(
                API_URL,
                headers=headers,
                data=data,
                files=files,
                timeout=10
            )
            response.raise_for_status()
            print(f"上传成功: {class_name} - {response.status_code}")

        except requests.exceptions.RequestException as e:
            print(f"网络错误: {str(e)}")
        except Exception as e:
            print(f"上传异常: {str(e)}")

    threading.Thread(target=_upload, daemon=True).start()


def send_mqtt_message(class_name, client):
    """
    向指定的 MQTT 主题发送检测日志
    """

    # 获取当前日期时间并格式化
    current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    # 构造 JSON 消息
    message = {
        "ICID": 1,  # 摄像头设备 ID
        "Datetime": current_time,  # 当前日期时间
        "Level1Class": 0,  # 事件分类
        "Log": ""  # 日志内容
    }

    if class_name == "person":
        message["Level1Class"] = 2
        message["Log"] = "发现人物入侵"
    elif class_name == "fire":
        message["Level1Class"] = 3
        message["Log"] = "发现火灾"
    else:
        message["Level1Class"] = 1
        message["Log"] = "发现病虫害"

    # 将字典转换为 JSON 字符串
    json_message = json.dumps(message)

    try:
        # 发布消息
        client.publish(topic, json_message)

        print(f"检测日志消息已发送到主题 '{topic}': {json_message}")

    except Exception as e:
        print(f"检测日志发送消息时出错: {e}")


def send_mqtt_ICHB(client):
    """
    定时向指定MQTT主题发送心跳检测的函数
    """

    def sender():
        # 构造 JSON 消息
        msg = {
            "ICID": 1,
            "HB": 1
        }

        # 将字典转换为 JSON 字符串
        json_msg = json.dumps(msg)

        try:
            # 发布消息
            client.publish(ICHB_topic, json_msg)

            print(f"心跳检测消息已发送到主题 '{ICHB_topic}': {json_msg}")

        except Exception as e:
            print(f"心跳检测发送消息时出错: {e}")

    def timed_sender():
        while True:
            sender()
            # 定时发送消息
            time.sleep(5)

    # 创建并启动线程
    send_thread = threading.Thread(target=timed_sender)
    send_thread.daemon = True  # 设置为守护线程，当主线程退出时子线程也会退出
    send_thread.start()


def main():
    last_upload_time = defaultdict(datetime.now)  # 记录各类别上次上传时间

    # 创建 MQTT 客户端实例
    client = mqtt.Client()
    # 连接至 MQTT 代理服务器
    client.connect(broker, port, keepalive)

    # 创建并启动心跳检测发送线程
    send_mqtt_ICHB(client)

    # 初始化模型
    model = YOLO(MODEL_PATH)
    class_names = model.names  # 获取类别名称字典

    # 初始化摄像头
    cap = cv2.VideoCapture(CAMERA_INDEX)
    cap.set(cv2.CAP_PROP_FRAME_WIDTH, CAMERA_RESOLUTION[0])
    cap.set(cv2.CAP_PROP_FRAME_HEIGHT, CAMERA_RESOLUTION[1])
    actual_fps = cap.get(cv2.CAP_PROP_FPS)  # 获取摄像头实际帧率
    if actual_fps <= 0:
        actual_fps = 30  # 默认值（部分摄像头无法获取帧率）

    # 计算检测间隔（确保实际处理帧率<=TARGET_FPS）
    detect_interval = max(1, int(actual_fps / TARGET_FPS))

    # 初始化FFmpeg推流进程
    ffmpeg_process = init_ffmpeg_stream()

    frame_counter = 0
    try:
        while cap.isOpened():
            ret, frame = cap.read()
            if not ret:
                print("摄像头读取失败")
                break

            frame_counter += 1

            # 跳帧处理：只有符合间隔的帧才处理并推流
            if frame_counter % detect_interval == 0:
                # 执行检测
                annotated_frame, detections = process_frame(frame, model)

                # 推流标注后的帧
                ffmpeg_process.stdin.write(annotated_frame.tobytes())

                # 处理检测结果（上传API）
                for box in detections:
                    class_id = int(box.cls[0])
                    class_name = class_names[class_id]
                    current_time = datetime.now()

                    # 检查5分钟内是否已上传过同类
                    # if (current_time - last_upload_time[class_name]) > timedelta(minutes=5):
                    if (current_time - last_upload_time[class_name]) > timedelta(seconds=10):
                        print("上传：", class_name)
                        upload_image(annotated_frame, class_name)
                        last_upload_time[class_name] = current_time

                        # 推送 MQTT 日志
                        send_mqtt_message(class_name, client)

            else:
                # 直接跳过不推流（避免原始帧与标注帧交替出现闪烁）
                continue

            # 控制循环频率（根据实际摄像头帧率）
            time.sleep(1 / actual_fps)

    except KeyboardInterrupt:
        print("用户中断")
    finally:
        # 释放资源
        cap.release()
        if ffmpeg_process:
            ffmpeg_process.stdin.close()
            ffmpeg_process.wait()
        print("资源已释放")
        # 断开MQTT连接
        client.disconnect()


if __name__ == "__main__":
    main()
