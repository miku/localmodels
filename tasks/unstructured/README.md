# Unstructured

Sample of 10000 unstructured strings, vary the prompt and model. Bibliographic
records. Try to output JSON.

With curl:

```shell
$ curl -sL 'localhost:11434/api/generate?format=json' -d '{"model": "mistral", "prompt": "Parse the following reference string into JSON: Amis, M. (2001, March 17). A rough trade : The Guardian. Retrieved from The Guardian: http:// www.theguardian.com/books/2001/mar/17/society.martinamis1"}'
```

![](622328.gif)

## Batch parse refs

```
"Przez integrację społeczną rozumiem tu procesy jednoczenia elementów i części składowych w jedną całość, które obejmują te wszystkie interakcje między składowymi częściami czy grupami, które prowadzą do ich powiązania, kooperacji, koegzystencji, dostosowania się czy do rozwiązywania konfliktów. Patrz: J. Turowski, Socjologia. Małe struktury społeczne, KUL, Lublin 1993, s. 129."
"CDU-Fraktion im Landtag von Sachsen-Anhalt (2017): Abgrenzen statt Ausgrenzen, Positionspapier vom 18.1.2017."
"Vega Ramos, María José. \"Malos saberes y censuras menores en el sigloX VI.\" Losm alos saberes.E d. Folke Gernert. Toulouse: Presses Universitaires du Midi, 2016. 13-28."
"Igel, Ulrike; Brähler, Elmar;"
"________________ La Revolución en Oaxaca (1900-1930), México, Consejo Nacional para la Cultura y las Artes,1993."
"Heitmeyer, Wilhelm 2018: Autoritäre Versuchungen. Berlin."
"(Y? @@@@@@@@@@@@@@@@@@@@@@@(M? @H ?J@@@@@@@@@@@@@@@@@@@@@@@H 5? O&@@@@@@@@@@@@@@@@@@@@@@@? ?7@@@@@@@@@@@@@@@@@@@@@@@@@@H?"
"Kuipers, Tom/ Hoeven van der, Jeffrey (2009): Insight into digital preservation of research output in Europe. Survey Report. PARSE.Insight D3.4. http://libereurope.eu/wp-content/uploads/2010/01/PARSE.Insight.-Deliverable- D3.4-Survey-Report.-of-research-output-Europe-Title-of-Deliverable-Survey-Report.pdf [Zugriff: 17.04.2018]."
"Marshall, Alfred (1920): Principles of Economics, London: MacMillan. References Martin, R. (1997): Regional Unemployment Disparities and their Dynamics, Regional Studies, 31: 237-252."
"OLG München vom 13.07.2000, Az.: 6 U 1791/00, Gewinnspiel, WRP 2000, 1321- 1323"
"Nguyen, Dan Thy: »Koalition der Freien Hamburg«, in: Kul- turpolitische Mitteilungen, Heft 145 (II/2014), S. 60-63"
"Parsons, Talcott (1958a): »The Professions and Social Struc- ture«. In: Ders., Essays in Sociological Theory, 2. Aufl. Glencoe: Free Press, S. 39-49."
"Operation Deliberate Force;"
"Suprunovskyi, A. (2015). Doctrinal problems of the study of migration processes."
"Otto Neurath, \" Wissenschaftliche Weltauffassung. Der Wiener Kreis,\" in Gesammelte philosophische und methodologische Schriften, ed. Rudolf Haller and Heiner Rutte (Vienna: Hölder-Pichler-Tempsky, 1981), Bd. 1, 304. Neurath sharp ened the intentions of the Vienna circle even more when he noticed in a review of Rudolf Carnap's \" Logischem Aufbau der Welt\": \"Antimetaphysicians (strengthen) the force of the proletariat\" (Neurath, Bd. 1, 338)."
"Land Brandenburg (2002): Ergänzung zur Programmplanung zum OP Brandenburg Förderperiode 2000 -2006."
"J. res.: fundam. care. online 2017. abr./jun. 9(2): 309-314"
```

Parsing 20 raw refs into json takes about

```
$ time shuf -n 20 refs.json | python loop.py > refs_parsed.json
```

Some example output:

```json
{
  "input": "\"Bundesagentur für Arbeit (2007b): Dokumentation der vorgenommenen Änderungen und Korrekturen bei der Bereitstellung der SGB II-Kennzahlen für interregionale Vergleiche, Stand: 19.06.2007, Nürnberg.\"\n",
  "parsed": {
    "Source": "Bundesagentur für Arbeit (2007b)",
    "DocumentTitle": "Dokumentation der vorgenommenen Änderungen und Korrekturen bei der Bereitstellung der SGB II-Kennzahlen für interregionale Vergleiche",
    "Date": "19.06.2007",
    "Location": "Nürnberg"
  }
}
{
  "input": "\"Taiwan So -tokufu gaiji-bu (1943), Kakyo -keizai jijyo -, Taiwan So -to -kufu gaiji cho -sa No. 133 [Economic situations of overseas Chinese, Foreign affairs survey No. 133. Office of the Governor-in-General, Taiwan]. Taipei: Ko -myo -sha sho -kai.\"\n",
  "parsed": {
    "title": "Taiwan So -tokufu gaiji-bu (1943)",
    "publisher": "Office of the Governor-in-General, Taiwan",
    "location": "Taipei: Ko -myo -sha sho -kai"
  }
}
{
  "input": "\"Bremer Kaufmann.\"\n",
  "parsed": {
    "first_name": "Bremer",
    "last_name": "Kaufmann."
  }
}
{
  "input": "\"Friedan, Betty: Der Weiblichkeitswahn. Ein vehementer Protest gegen das Wunschbild von der Frau. Hamburg 1966\"\n",
  "parsed": {
    "author": "Betty Friedan",
    "title": "Der Weiblichkeitswahn. Ein vehementer Protest gegen das Wunschbild von der Frau.",
    "location": "Hamburg",
    "year": 1966
  }
}
```
