from ocr import ocr
import numpy as np
import cv2 as cv
from flask import Flask, jsonify, request


def single_pic_proc(img):
    img = cv.cvtColor(img, cv.COLOR_BGR2RGB)
    _, img = cv.threshold(img, 127, 255, cv.THRESH_TOZERO)
    image = np.array(img)
    result, image_framed = ocr(image)
    ans = ''
    for key in result:
        result[key][1].replace("!", "1")
        print(result[key])
        # if result[key][1].isalnum() and len(result[key][1]) > 3:
        if 1:
            ans += result[key][1]
    return ans, image_framed


app = Flask(__name__)
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'JPG', 'PNG', 'gif', 'GIF'])


def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1] in ALLOWED_EXTENSIONS


# 上传文件
@app.route('/upload', methods=['POST'], strict_slashes=False)
def api_upload():
    f = request.files['picture']
    img = f.read()
    img = cv.imdecode(np.frombuffer(img, np.uint8), cv.IMREAD_COLOR)
    if f and allowed_file(f.filename):
        result, image_framed = single_pic_proc(img)
        return {
            "status": 200,
            "msg": result,
        }
    else:
        return {
            "status": 1001,
            "msg": "上传失败",
        }


@app.route('/test', methods=['GET'])
def demo():
    return {
        "status": 200,
        "msg": "123456"
    }


if __name__ == '__main__':
    app.run(debug=True)
