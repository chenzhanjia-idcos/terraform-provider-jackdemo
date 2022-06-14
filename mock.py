from flask import request, Flask, jsonify, abort

app = Flask(__name__)
app.config['JSON_AS_ASCII'] = False
recognize_info = {}

@app.route('/', methods=['GET'])
def init_api():
    return "success", 200


@app.route('/create', methods=['POST'])
def post_Data():
    recognize_info["instance_name"] = request.json['instance_name']
    recognize_info["disk_size"] = request.json['disk_size']
    recognize_info["tags"] = request.json['tags']
    recognize_info["id"] = "jackthebest1"
    return jsonify(recognize_info), {"create_success": 200}

@app.route('/get', methods=['GET'])
def get_Data():
    if request.args.get("id") == "jackthebest1":
        return jsonify(recognize_info), 200
    abort(400)

@app.route('/update', methods=['PUT'])
def update_Data():
    if request.args.get("id") != "jackthebest1":
        abort(400)
    if request.json['instance_name']:
        recognize_info["instance_name"] = request.json['instance_name']
    if request.json['disk_size']:
        recognize_info["disk_size"] = request.json['disk_size']
    if request.json['tags']:
        recognize_info["tags"] = request.json['tags']
    return jsonify(recognize_info), 200

@app.route('/delete', methods=['DELETE'])
def delete_Data():
    if request.args.get("id") == "jackthebest1":
        recognize_info = {}
        return jsonify(recognize_info), 200
    abort(400)

@app.route("/data_source", methods=["GET"])
def data_source():
    if request.args.get("name") == "ecs":
        rep = {"name": "ecs",
               "tags" :"test",
               "instance_type": "normal",
               "id" : "ecs_id",
               }
        return jsonify(rep), 200
    return jsonify({}), 200



if __name__ == '__main__':
    app.run(debug=False, host='0.0.0.0', port=8888)


# 安装依赖 pip3 install flask
# 启动服务： python3 mock.py
