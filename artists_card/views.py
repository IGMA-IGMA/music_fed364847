from django.shortcuts import render, get_object_or_404
<<<<<<< HEAD
from .models import Artist_page
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt

def artist_detail(request, artist_id):
    artist = get_object_or_404(Artist_page, id=artist_id)
    return render(request, 'artist_page.html', {'artist': artist})
=======
from .models import Artist_page, Artist_song
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt
from django.db.models import Q

def artist_detail(request, artist_id):
    artist = get_object_or_404(Artist_page, id=artist_id)
    songs = artist.songs.all().order_by('-plays')
    if songs.exists():
        total_plays_in_months = sum(song.plays for song in songs)/songs.count()
    else:
        total_plays_in_months = 0
    context = {
        'artist': artist,
        'songs': songs,
        'songs_count': songs.count(),
        'total_plays_in_months': int(total_plays_in_months),
    }
    return render(request, 'artist_page.html', context)
>>>>>>> 8db0ba0dbecf88af941a3ee9bff345731e3e4735

@csrf_exempt
def submit_data(request):
    if request.method == 'POST':
        username = request.POST.get('username')
        password = request.POST.get('password')
        email = request.POST.get('email')
        return JsonResponse({
            'status': 'success',
            'message': 'Данные получены и обработаны'
        })
<<<<<<< HEAD
    return render(request, 'auth_form.html')
=======
    return render(request, 'auth_form.html')

def artists_all(request):
    artists = Artist_page.objects.all()
    context = {
        'artists': artists,
        'total_count': artists.count(),
    }
    return render(request, 'mainpage.html', context)

def global_search(request):
    search_query = request.GET.get('q','').strip()
    artists_result = []
    songs_result = []
    if search_query:
        artists_result = Artist_page.objects.filter(
            name__icontains=search_query
        ).order_by('-listeners')
        songs_result = Artist_song.objects.filter(
            Q(title__icontains=search_query) |
            Q(artist__name__icontains=search_query)
        ).order_by('-plays')
    context = {
        'search_query': search_query,
        'artists': artists_result,
        'songs': songs_result,
        'artists_count': len(artists_result),
        'songs_count': len(songs_result),
        'total_results': len(artists_result) + len(songs_result),
    }
    return render(request, 'search_results.html', context)
>>>>>>> 8db0ba0dbecf88af941a3ee9bff345731e3e4735
