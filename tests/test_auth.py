import requests, os
from random_username.generate import generate_username

host = f"http://{os.getenv('API_HOST')}"

def signup():
    auth_data = generate_username(2)
    username = auth_data[0][:20]
    password = auth_data[1][:64]
    response = requests.post(f"{host}/am/signup", params={
        "username": username, "password": password
    })
    return username, password, response.json()

def signin(username, password):
    response = requests.post(f"{host}/am/signin", params={
        "username": username, "password": password
    })
    return response.json()
