# Otaniemen-Ruokalistat
Yksinkertainen Go ohjelma, joka hakee ISS:n nettisivustolta (https://ravintolapalvelut.iss.fi/espoon-tietokylä) Otaniemen lukion ruokalistan pyydettäessä. 
Idea varastettu toiselta telegram botilta joka ei ilmeisesti ole ainakaan tällä hetkellä käytössä.
## Setup
```
git clone https://github.com/Aapeli123/Otaniemen-Ruokalistat.git
cd Otaniemen-Ruokalistat
```
Hae telegram api-token ja aseta se muuttujaan token tiedostossa `./telegram/bot.go`
Jos et tarvitse JSON APIa kommentoi pois `main.go` tiedostosta käsky `go webapi.Init()`
Telegram botin saa pois kommentoimalla käskyn `go telegram.Init()`

```
go build .
sudo ./otaniemenruokalistat.tk
```
## JSON API
Ohjelman suorittaminen avaa portin 9999 ja aloittaa sinne webserverin. Sen route `/json` antaa viikon ruokalistan JSON muodossa
### Esimerkkivastaus:
```json
[
  {
    "paiva": "Maanantai 19.10.",
    "kotiruoka": "Broilerikiusausta L G Vs",
    "kasvisruoka": "Kasviskiusausta  L G Vs"
  },
  {
    "paiva": "Tiistai 20.10.",
    "kotiruoka": "Makkarakastiketta ja perunoita M C",
    "kasvisruoka": "Meksikolaista nyhtöhernepataa M G C VEG"
  },
  {
    "paiva": "Keskiviikko 21.10.",
    "kotiruoka": "Kebabkastiketta ja riisiä M G Vs C",
    "kasvisruoka": "Tofu currya ja riisiä L G Vs"
  },
  {
    "paiva": "Torstai 22.10.",
    "kotiruoka": "Lohikeittoa L G Ka",
    "kasvisruoka": "Espanjalaista kasviskeittoa L G Vs"
  },
  {
    "paiva": "Perjantai 23.10.",
    "kotiruoka": "Bataatti-soijavuokaa M G Vs VEG So",
    "kasvisruoka": "Kikherne-papupataa L G Vs So"
  }
]
```
## TODOs
Telegram botista Inline mode
