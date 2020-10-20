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

```
go build .
sudo ./otaniemenruokalistat.tk
```

## TODOs
Telegram botista Inline mode
