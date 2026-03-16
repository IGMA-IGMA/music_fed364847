from django.db import models
from django.urls import reverse

class Artist_page(models.Model):
    name = models.CharField(max_length = 200, null=False,verbose_name='Имя')
    listeners = models.IntegerField(default=0,verbose_name='Кол-во слушателей')
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