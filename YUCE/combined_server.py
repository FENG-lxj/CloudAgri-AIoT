# combined_server.py
import os
import sys
from multiprocessing import Process
import subprocess
import time


def run_production_app():
    """启动作物产量预测服务"""
    print("🚀 正在启动作物产量预测服务...")
    try:
        subprocess.run(
            [sys.executable, "productionapp.py"],
            check=True,
            stdout=sys.stdout,
            stderr=sys.stderr
        )
    except subprocess.CalledProcessError as e:
        print(f"❌ 产量预测服务异常退出: {e}")
    except KeyboardInterrupt:
        print("🛑 产量预测服务已停止")


def run_crop_recommender():
    """启动作物推荐服务"""
    print("🌱 正在启动作物推荐服务...")
    try:
        subprocess.run(
            [sys.executable, "crop_recommender.py"],
            check=True,
            stdout=sys.stdout,
            stderr=sys.stderr
        )
    except subprocess.CalledProcessError as e:
        print(f"❌ 作物推荐服务异常退出: {e}")
    except KeyboardInterrupt:
        print("🛑 作物推荐服务已停止")


def run_weather_service():
    """启动天气服务（tianqi.py）"""
    print("🌤️ 正在启动天气服务...")
    try:
        subprocess.run(
            [sys.executable, "tianqi.py"],
            check=True,
            stdout=sys.stdout,
            stderr=sys.stderr
        )
    except subprocess.CalledProcessError as e:
        print(f"❌ 天气服务异常退出: {e}")
    except KeyboardInterrupt:
        print("🛑 天气服务已停止")


def run_combined_server():
    """主启动函数"""
    print("\n" + "=" * 50)
    print("🌾 智能农业服务系统启动器")
    print("=" * 50)

    # 创建进程
    processes = [
        Process(target=run_production_app),     # 端口: 5000
        Process(target=run_crop_recommender),   # 端口: 5001
        Process(target=run_weather_service)     # 端口: 4999（或其他你在 tianqi.py 中定义的端口）
    ]

    try:
        # 启动所有服务
        for p in processes:
            p.start()
            time.sleep(1)  # 避免输出混乱

        print("\n✅ 服务状态:")
        print(f"  产量预测: http://localhost:5000")
        print(f"  作物推荐: http://localhost:5001")
        print(f"  天气服务: http://localhost:4999")  # 请根据你的 tianqi.py 的实际端口修改
        print("\n🔄 运行中... (Ctrl+C 停止所有服务)")

        # 监控进程
        while True:
            if not any(p.is_alive() for p in processes):
                print("⚠️ 检测到服务异常终止")
                break
            time.sleep(5)

    except KeyboardInterrupt:
        print("\n🛑 正在停止所有服务...")
        for p in processes:
            if p.is_alive():
                p.terminate()
                p.join(timeout=3)
        print("所有服务已安全停止")

    except Exception as e:
        print(f"❌ 启动器错误: {str(e)}")
        for p in processes:
            if p.is_alive():
                p.terminate()


if __name__ == '__main__':
    run_combined_server()