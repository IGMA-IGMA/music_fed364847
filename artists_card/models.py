from django.db import models
from django.urls import reverse
from django.core.exceptions import ValidationError
import os

def validate_mp3_file(value):
    ext = os.path.splitext(value.name)[1]
    valid_extensions = ['.mp3']
    if not ext.lower() in valid_extensions:
        raise ValidationError('Разрешены только MP3 файлы.')

class Artist_page(models.Model):
    name = models.CharField(max_length = 200, null=False,verbose_name='Имя')
    listeners = models.IntegerField(default=0,verbose_name='Кол-во слушателей в месяц')
    bio = models.CharField(max_length=258,verbose_name='Описание')
    photo = models.ImageField(
        upload_to='artists/', 
        blank=True, 
        null=True, 
        verbose_name='Фото',
        default='artists/media/artists/images.jpg'
    )

    class Meta:
        verbose_name = 'Артист'
        verbose_name_plural = 'Артисты'
        ordering = ['-listeners']

    def __str__(self):
        return self.name
    
    def get_absolute_url(self):
        return reverse('artist_detail', args=[str(self.id)])

class Artist_song(models.Model):
    artist = models.ForeignKey(Artist_page, on_delete=models.CASCADE, related_name='songs', verbose_name='Артист')
    title = models.CharField(max_length=200, verbose_name='Название песни')
    duration = models.CharField(max_length=10, blank=True, help_text='Формат: 3:45', verbose_name='Длительность')
    plays = models.IntegerField(default=0, verbose_name='Количество прослушиваний')
    audio_file = models.FileField(upload_to='songs/', blank=True, null=True, verbose_name='Аудиофайл', validators=[validate_mp3_file])
    class Meta:
        verbose_name = 'Песня'
        verbose_name_plural = 'Песни'
        ordering = ['-plays', 'title']
    def __str__(self):
        return f'{self.artist.name} - {self.title}'