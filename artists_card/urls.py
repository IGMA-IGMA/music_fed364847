from django.urls import path
from . import views

urlpatterns = [
    path('artist/<int:artist_id>/', views.artist_detail, name='artist_detail'),
    path('auth/', views.submit_data, name='auth'),
]