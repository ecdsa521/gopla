# gopla
Another ripla clone.

Usage:
---

    -a	Grab all hashes for search query
    -b	Only print best option
    -l	Grab all links for search query
    -s string
            Search for this show
    -v	Verbose operation
    -w	Generate batch output for wget downloads
  
  Examples:
  ---
  
  Find a show by name
  
    $ gopla -v -s "Miodowe Lata"
    Miodowe Lata
    Dwaj przyjaciele: Karol Krawczyk (Cezary Żak) - motorniczy tramwaju i Tadzio Norek (Artur Barciś) - z zawodu kanalarz - mieszkają wraz ze swoimi żonami w jednej kamienic w Warszawie, w robotniczej dzielnicy Wola. Są niepoprawnymi optymistami i ciągle wierzą, że zdobędą fortunę. A tymczasem mają różne szalone pomysły, z czego wynika wiele przezabawnych perypetii. Ale ich żony: Alinka Krawczyk (Agnieszka Piliszewska/Katarzyna Żak) i Danusia Norek (Dorota Chotecka-Pazura) wiedzą najlepiej, jak należy z nimi postępować i ściągnąć na ziemię...
    http://www.ipla.tv/kategoria/5007481 | null


  Get hashes for show's episodes
  
    $ gopla -s "Miodowe Lata" -a
    6df54484ee2f09a22ff3a8b46c44f328
    6a07992bc4f0879da697035dea1fcf3e
    8909aba012e0d294d24d5d0f6fe5fd8b
    8b5dcbef6f4982c8af9cc6536c5b8de6
    (...)

  Get information about hash
  
    $ gopla -v 6df54484ee2f09a22ff3a8b46c44f328
    Miodowe Lata - Mur czyli zemsta (00:40:33)
    Alina pracowała dorywczo w firmie budowlanej. Zamiast wynagrodzenia zażyczyła sobie, aby firma wymurowała jej w kuchni blat. Karolowi oczywiście nie podoba się takie rozwiązanie. Wolałby gotówkę. Gdy przychodzi fachowiec wykonać zlecenie, Karol się z nim kłóci i obraża. Gdy z kolei u Krawczyków zjawia się Tadzio, znajduje Karola zamurowanego w sypialni. Co się stało???

    
