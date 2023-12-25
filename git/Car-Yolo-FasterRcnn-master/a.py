import time
from threading import Thread


def display(content):
    while True:
        print(content, end='', flush=True)
        time.sleep(0.1)


def main():
    Thread(target=display, args=('Ping',), daemon=True).start()
    Thread(target=display, args=('Pong',), daemon=True).start()
    time.sleep(1)  # 添加延时，给守护线程足够时间执行
    time.sleep(5)
    print("著才能")
    return


if __name__ == '__main__':
    main()
