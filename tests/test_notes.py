import requests, os, json
import test_auth

host = f"http://{os.getenv('API_HOST')}"

def get(username, token):
    response = requests.post(f"{host}/notebook/get", params={
        "username": username, "token": token
    })
    print(response.text)
    return response.json()

def add(username, token):
    response = requests.post(f"{host}/notebook/add", params={
        "username": username, "token": token
    })
    print(response.text)
    return response.json()

def update(username, token, data):
    response = requests.post(f"{host}/notebook/update", params={
        "username": username, "token": token
    }, json={
        "id": int(data["id"]), "title": test_auth.generate_username(1)[0],
        "text": test_auth.generate_username(1)[0], "checked": True
    })
    print(response.text)
    return response.json()

def delete(username, token, data):
    response = requests.post(f"{host}/notebook/delete", params={
        "username": username, "token": token
    }, json={
        "id": int(data["id"])
    })
    print(response.text)
    return response.json()
