import test_notes
import test_auth
from pprint import pprint


username, password, response = test_auth.signup()
response = test_auth.signin(username, password)
token = response["token"]

response = test_notes.get(username, token)
pprint(response)
for _ in range(3):
    response = test_notes.add(username, token)
    pprint(response)
response = test_notes.get(username, token)
pprint(response)
response = test_notes.update(username, token, {"id": int(input("[UPDATE]===> "))})
response = test_notes.get(username, token)
pprint(response)
test_notes.delete(username, token, {"id": int(input("[DELETE]===> "))})
response = test_notes.get(username, token)
pprint(response)
