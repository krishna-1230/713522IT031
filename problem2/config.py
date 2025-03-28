from typing import Dict

# Base URL for the social media platform API
BASE_URL = "http://20.244.56.144/test"

# Authentication token
AUTH_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzQzMTU1MDM1LCJpYXQiOjE3NDMxNTQ3MzUsImlzcyI6IkFmZm9yZG1lZCIsImp0aSI6ImY5YmY5YWZjLTBkYmEtNDNjZS04MTViLTJjNTlhZDdiOGNjYiIsInN1YiI6ImtyaXNoc3B5azEyMzBAZ21haWwuY29tIn0sImNvbXBhbnlOYW1lIjoiZ29NYXJ0IiwiY2xpZW50SUQiOiJmOWJmOWFmYy0wZGJhLTQzY2UtODE1Yi0yYzU5YWQ3YjhjY2IiLCJjbGllbnRTZWNyZXQiOiJPVmp1RHV5R2JwQWZoSGJ4Iiwib3duZXJOYW1lIjoiUmFodWwiLCJvd25lckVtYWlsIjoia3Jpc2hzcHlrMTIzMEBnbWFpbC5jb20iLCJyb2xsTm8iOiI3MTM1MjJJVDAzMSJ9.mDznIxHTSdSGUktKLmtY1FgzShQhmyIBzqA84tuqtxk"

# Cache settings (in seconds)
CACHE_TTL = 60  # Cache time to live

# API endpoints
ENDPOINTS = {
    "users": f"{BASE_URL}/users",
    "user_posts": f"{BASE_URL}/users/{{user_id}}/posts",
    "post_comments": f"{BASE_URL}/posts/{{post_id}}/comments"
}

# Response models
class User(Dict):
    id: str
    name: str

class Post(Dict):
    id: int
    userid: int
    content: str

class Comment(Dict):
    id: int
    postid: int
    content: str 