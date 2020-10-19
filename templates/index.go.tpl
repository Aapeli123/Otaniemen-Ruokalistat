<html>
<head>
    <title> Otaniemen Lukion Ruokalistat </title>
    <style>
    body {
        text-align: center;
    }
    </style>
</head>
<body>
    <h1>Ruokalista:</h1>
    {{range .Paivat}}
        <h2> {{.Viikonpäivä}} </h2>
        <h3> {{.Perus}} </h3>
        <h3> {{.Veg}} </h3>
    {{end}}
    <p> 
    VL = Vähälaktoosinen, L = Laktoositon, M = Maidoton, G = Gluteeniton
    </p>
</body>
<html>