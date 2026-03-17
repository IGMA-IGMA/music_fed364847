from django.contrib import admin
from .models import Artist_page

@admin.register(Artist_page)
class ArtistAdmin(admin.ModelAdmin):
    list_display = ('name', 'listeners')
    list_filter = ('listeners',)
    search_fields = ('name', 'bio')
    list_editable = ('listeners',)