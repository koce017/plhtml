U trenutnoj implementaciji se tabela simbola koristi kako za provere tokom semantičke analize, tako i za čuvanje
vrednosti promenljvih u fazi interpretacije. Tu je postojao jedan problem sa variable shadowing-om.

```html
<var class="integer">a</var>
<div data-if="true">
    <output>a + `\n`</output>
    <var class="integer">a</var>
</div>
```

Pošto se tabela simbola popunjava u fazi semantičke analize, u trenutku interpratiranja `output` naredbe u tabeli
simbola se već nalaze dva `a`: jedno u root opsegu, a drugo u unutrašnjem opsegu. Lookup je umesto `a` iz spoljašnjeg
vraćao `a` iz unutrašnjeg opsega, koje u tom trenutku ima `nil` vrednost, s obzirom da se inicijalizacija nulom izvršava
u fazi interpretacije.

Ovo je rešeno otprilike tako što se identifikatorima u stablu tokom semantičke analize doda broj linije u kojoj su
deklarisani, i da ne bi bilo zabune to se koristi za identifikovanje. Pri referenciranju promenljivih u kodu se ne piše
broj linije deklaracije, pa u lookup metodi postoji i pretraživanje samo po imenu identifikatora, koje nazad vraća
dopunjeni naziv. (fd967b4e)