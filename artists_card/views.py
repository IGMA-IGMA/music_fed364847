from django.shortcuts import render, get_object_or_404
from .models import Artist_page

def artist_detail(request, artist_id):
    artist = get_object_or_404(Artist_page, id=artist_id)
    return render(request, 'artist_page.html', {'artist': artist})