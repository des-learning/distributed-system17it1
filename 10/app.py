from flask import Flask
import subprocess

app = Flask(__name__)

@app.route('/')
def index():
    hostname = subprocess.run(['hostname'], capture_output=True).stdout.strip().decode()
    return "Hello from {}".format(hostname)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)