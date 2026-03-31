from django.urls import path
from . import views

urlpatterns = [
    path('artist/<int:artist_id>/', views.artist_detail, name='artist_detail'),
    path('auth/', views.submit_data, name='auth'),
<<<<<<< HEAD
=======
    path('', views.artists_all, name ='home'),
    path('search/', views.global_search, name='global_search'),
>>>>>>> 8db0ba0dbecf88af941a3ee9bff345731e3e4735
]