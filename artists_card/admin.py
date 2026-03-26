from django.contrib import admin
from .models import Artist_page, Artist_song

@admin.register(Artist_page)
class ArtistAdmin(admin.ModelAdmin):
    list_display = ('name', 'listeners')
    list_filter = ('listeners',)
    search_fields = ('name', 'bio')
    list_editable = ('listeners',)

@admin.register(Artist_song)
class SongAdmin(admin.ModelAdmin):
    list_display = ('title', 'artist', 'duration', 'plays')
    list_filter = ('artist',)
    search_fields = ('title', 'artist__name')
    list_editable = ('plays',) 
    list_select_related = ('artist',)