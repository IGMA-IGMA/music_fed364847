from django.shortcuts import render, get_object_or_404
from .models import Artist_page
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt

def artist_detail(request, artist_id):
    artist = get_object_or_404(Artist_page, id=artist_id)
    return render(request, 'artist_page.html', {'artist': artist})

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
    return render(request, 'auth_form.html')