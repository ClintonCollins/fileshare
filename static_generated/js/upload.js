var X=Object.defineProperty;var Q=(u,r,n)=>r in u?X(u,r,{enumerable:!0,configurable:!0,writable:!0,value:n}):u[r]=n;var d=(u,r,n)=>(Q(u,typeof r!="symbol"?r+"":r,n),n);import{M as _}from"./MediaPreviewModal.js";/**
 * filesize
 *
 * @copyright 2022 Jason Mulligan <jason.mulligan@avoidwork.com>
 * @license BSD-3-Clause
 * @version 10.0.6
 */const H="array",ee="bit",K="bits",te="byte",W="bytes",j="",ne="exponent",re="function",N="iec",ie="Invalid number",ue="Invalid rounding method",A="jedec",se="object",Z=".",oe="round",le="s",ae="kbit",me="kB",de=" ",ce="string",fe="0",G={symbol:{iec:{bits:["bit","Kibit","Mibit","Gibit","Tibit","Pibit","Eibit","Zibit","Yibit"],bytes:["B","KiB","MiB","GiB","TiB","PiB","EiB","ZiB","YiB"]},jedec:{bits:["bit","Kbit","Mbit","Gbit","Tbit","Pbit","Ebit","Zbit","Ybit"],bytes:["B","KB","MB","GB","TB","PB","EB","ZB","YB"]}},fullform:{iec:["","kibi","mebi","gibi","tebi","pebi","exbi","zebi","yobi"],jedec:["","kilo","mega","giga","tera","peta","exa","zetta","yotta"]}};function Y(u,{bits:r=!1,pad:n=!1,base:o=-1,round:f=2,locale:p=j,localeOptions:y={},separator:E=j,spacer:v=de,symbols:M={},standard:h=j,output:x=ce,fullform:l=!1,fullforms:g=[],exponent:C=-1,roundingMethod:F=oe,precision:T=0}={}){let s=C,e=Number(u),t=[],i=0,a=j;o===-1&&h.length===0?(o=10,h=A):o===-1&&h.length>0?(h=h===N?N:A,o=h===N?2:10):(o=o===2?2:10,h=o===10||h===A?A:N);const m=o===10?1e3:1024,w=l===!0,k=e<0,b=Math[F];if(typeof u!="bigint"&&isNaN(u))throw new TypeError(ie);if(typeof b!==re)throw new TypeError(ue);if(k&&(e=-e),(s===-1||isNaN(s))&&(s=Math.floor(Math.log(e)/Math.log(m)),s<0&&(s=0)),s>8&&(T>0&&(T+=8-s),s=8),x===ne)return s;if(e===0)t[0]=0,a=t[1]=G.symbol[h][r?K:W][s];else{i=e/(o===2?Math.pow(2,s*10):Math.pow(1e3,s)),r&&(i=i*8,i>=m&&s<8&&(i=i/m,s++));const S=Math.pow(10,s>0?f:0);t[0]=b(i*S)/S,t[0]===m&&s<8&&C===-1&&(t[0]=1,s++),a=t[1]=o===10&&s===1?r?ae:me:G.symbol[h][r?K:W][s]}if(k&&(t[0]=-t[0]),T>0&&(t[0]=t[0].toPrecision(T)),t[1]=M[t[1]]||t[1],p===!0?t[0]=t[0].toLocaleString():p.length>0?t[0]=t[0].toLocaleString(p,y):E.length>0&&(t[0]=t[0].toString().replace(Z,E)),n&&Number.isInteger(t[0])===!1&&f>0){const S=E||Z,R=t[0].toString().split(S),D=R[1]||j,I=D.length,P=f-I;t[0]=`${R[0]}${S}${D.padEnd(I+P,fe)}`}return w&&(t[1]=g[s]?g[s]:G.fullform[h][s]+(r?ee:te)+(t[0]===1?j:le)),x===H?t:x===se?{value:t[0],symbol:t[1],exponent:s,unit:a}:t.join(v)}var V={},he={get exports(){return V},set exports(u){V=u}};(function(u){(function(){var r={y:function(e){return e===1?"χρόνος":"χρόνια"},mo:function(e){return e===1?"μήνας":"μήνες"},w:function(e){return e===1?"εβδομάδα":"εβδομάδες"},d:function(e){return e===1?"μέρα":"μέρες"},h:function(e){return e===1?"ώρα":"ώρες"},m:function(e){return e===1?"λεπτό":"λεπτά"},s:function(e){return e===1?"δευτερόλεπτο":"δευτερόλεπτα"},ms:function(e){return(e===1?"χιλιοστό":"χιλιοστά")+" του δευτερολέπτου"},decimal:","},n=["۰","١","٢","٣","٤","٥","٦","٧","٨","٩"],o={af:{y:"jaar",mo:function(e){return"maand"+(e===1?"":"e")},w:function(e){return e===1?"week":"weke"},d:function(e){return e===1?"dag":"dae"},h:function(e){return e===1?"uur":"ure"},m:function(e){return e===1?"minuut":"minute"},s:function(e){return"sekonde"+(e===1?"":"s")},ms:function(e){return"millisekonde"+(e===1?"":"s")},decimal:","},ar:{y:function(e){return["سنة","سنتان","سنوات"][h(e)]},mo:function(e){return["شهر","شهران","أشهر"][h(e)]},w:function(e){return["أسبوع","أسبوعين","أسابيع"][h(e)]},d:function(e){return["يوم","يومين","أيام"][h(e)]},h:function(e){return["ساعة","ساعتين","ساعات"][h(e)]},m:function(e){return["دقيقة","دقيقتان","دقائق"][h(e)]},s:function(e){return["ثانية","ثانيتان","ثواني"][h(e)]},ms:function(e){return["جزء من الثانية","جزآن من الثانية","أجزاء من الثانية"][h(e)]},decimal:",",delimiter:" و ",_formatCount:function(e,t){for(var i=M(n,{".":t}),a=e.toString().split(""),m=0;m<a.length;m++){var w=a[m];s(i,w)&&(a[m]=i[w])}return a.join("")}},bg:{y:function(e){return["години","година","години"][l(e)]},mo:function(e){return["месеца","месец","месеца"][l(e)]},w:function(e){return["седмици","седмица","седмици"][l(e)]},d:function(e){return["дни","ден","дни"][l(e)]},h:function(e){return["часа","час","часа"][l(e)]},m:function(e){return["минути","минута","минути"][l(e)]},s:function(e){return["секунди","секунда","секунди"][l(e)]},ms:function(e){return["милисекунди","милисекунда","милисекунди"][l(e)]},decimal:","},bn:{y:"বছর",mo:"মাস",w:"সপ্তাহ",d:"দিন",h:"ঘন্টা",m:"মিনিট",s:"সেকেন্ড",ms:"মিলিসেকেন্ড"},ca:{y:function(e){return"any"+(e===1?"":"s")},mo:function(e){return"mes"+(e===1?"":"os")},w:function(e){return"setman"+(e===1?"a":"es")},d:function(e){return"di"+(e===1?"a":"es")},h:function(e){return"hor"+(e===1?"a":"es")},m:function(e){return"minut"+(e===1?"":"s")},s:function(e){return"segon"+(e===1?"":"s")},ms:function(e){return"milisegon"+(e===1?"":"s")},decimal:","},cs:{y:function(e){return["rok","roku","roky","let"][g(e)]},mo:function(e){return["měsíc","měsíce","měsíce","měsíců"][g(e)]},w:function(e){return["týden","týdne","týdny","týdnů"][g(e)]},d:function(e){return["den","dne","dny","dní"][g(e)]},h:function(e){return["hodina","hodiny","hodiny","hodin"][g(e)]},m:function(e){return["minuta","minuty","minuty","minut"][g(e)]},s:function(e){return["sekunda","sekundy","sekundy","sekund"][g(e)]},ms:function(e){return["milisekunda","milisekundy","milisekundy","milisekund"][g(e)]},decimal:","},cy:{y:"flwyddyn",mo:"mis",w:"wythnos",d:"diwrnod",h:"awr",m:"munud",s:"eiliad",ms:"milieiliad"},da:{y:"år",mo:function(e){return"måned"+(e===1?"":"er")},w:function(e){return"uge"+(e===1?"":"r")},d:function(e){return"dag"+(e===1?"":"e")},h:function(e){return"time"+(e===1?"":"r")},m:function(e){return"minut"+(e===1?"":"ter")},s:function(e){return"sekund"+(e===1?"":"er")},ms:function(e){return"millisekund"+(e===1?"":"er")},decimal:","},de:{y:function(e){return"Jahr"+(e===1?"":"e")},mo:function(e){return"Monat"+(e===1?"":"e")},w:function(e){return"Woche"+(e===1?"":"n")},d:function(e){return"Tag"+(e===1?"":"e")},h:function(e){return"Stunde"+(e===1?"":"n")},m:function(e){return"Minute"+(e===1?"":"n")},s:function(e){return"Sekunde"+(e===1?"":"n")},ms:function(e){return"Millisekunde"+(e===1?"":"n")},decimal:","},el:r,en:{y:function(e){return"year"+(e===1?"":"s")},mo:function(e){return"month"+(e===1?"":"s")},w:function(e){return"week"+(e===1?"":"s")},d:function(e){return"day"+(e===1?"":"s")},h:function(e){return"hour"+(e===1?"":"s")},m:function(e){return"minute"+(e===1?"":"s")},s:function(e){return"second"+(e===1?"":"s")},ms:function(e){return"millisecond"+(e===1?"":"s")},decimal:"."},eo:{y:function(e){return"jaro"+(e===1?"":"j")},mo:function(e){return"monato"+(e===1?"":"j")},w:function(e){return"semajno"+(e===1?"":"j")},d:function(e){return"tago"+(e===1?"":"j")},h:function(e){return"horo"+(e===1?"":"j")},m:function(e){return"minuto"+(e===1?"":"j")},s:function(e){return"sekundo"+(e===1?"":"j")},ms:function(e){return"milisekundo"+(e===1?"":"j")},decimal:","},es:{y:function(e){return"año"+(e===1?"":"s")},mo:function(e){return"mes"+(e===1?"":"es")},w:function(e){return"semana"+(e===1?"":"s")},d:function(e){return"día"+(e===1?"":"s")},h:function(e){return"hora"+(e===1?"":"s")},m:function(e){return"minuto"+(e===1?"":"s")},s:function(e){return"segundo"+(e===1?"":"s")},ms:function(e){return"milisegundo"+(e===1?"":"s")},decimal:","},et:{y:function(e){return"aasta"+(e===1?"":"t")},mo:function(e){return"kuu"+(e===1?"":"d")},w:function(e){return"nädal"+(e===1?"":"at")},d:function(e){return"päev"+(e===1?"":"a")},h:function(e){return"tund"+(e===1?"":"i")},m:function(e){return"minut"+(e===1?"":"it")},s:function(e){return"sekund"+(e===1?"":"it")},ms:function(e){return"millisekund"+(e===1?"":"it")},decimal:","},eu:{y:"urte",mo:"hilabete",w:"aste",d:"egun",h:"ordu",m:"minutu",s:"segundo",ms:"milisegundo",decimal:","},fa:{y:"سال",mo:"ماه",w:"هفته",d:"روز",h:"ساعت",m:"دقیقه",s:"ثانیه",ms:"میلی ثانیه",decimal:"."},fi:{y:function(e){return e===1?"vuosi":"vuotta"},mo:function(e){return e===1?"kuukausi":"kuukautta"},w:function(e){return"viikko"+(e===1?"":"a")},d:function(e){return"päivä"+(e===1?"":"ä")},h:function(e){return"tunti"+(e===1?"":"a")},m:function(e){return"minuutti"+(e===1?"":"a")},s:function(e){return"sekunti"+(e===1?"":"a")},ms:function(e){return"millisekunti"+(e===1?"":"a")},decimal:","},fo:{y:"ár",mo:function(e){return e===1?"mánaður":"mánaðir"},w:function(e){return e===1?"vika":"vikur"},d:function(e){return e===1?"dagur":"dagar"},h:function(e){return e===1?"tími":"tímar"},m:function(e){return e===1?"minuttur":"minuttir"},s:"sekund",ms:"millisekund",decimal:","},fr:{y:function(e){return"an"+(e>=2?"s":"")},mo:"mois",w:function(e){return"semaine"+(e>=2?"s":"")},d:function(e){return"jour"+(e>=2?"s":"")},h:function(e){return"heure"+(e>=2?"s":"")},m:function(e){return"minute"+(e>=2?"s":"")},s:function(e){return"seconde"+(e>=2?"s":"")},ms:function(e){return"milliseconde"+(e>=2?"s":"")},decimal:","},gr:r,he:{y:function(e){return e===1?"שנה":"שנים"},mo:function(e){return e===1?"חודש":"חודשים"},w:function(e){return e===1?"שבוע":"שבועות"},d:function(e){return e===1?"יום":"ימים"},h:function(e){return e===1?"שעה":"שעות"},m:function(e){return e===1?"דקה":"דקות"},s:function(e){return e===1?"שניה":"שניות"},ms:function(e){return e===1?"מילישנייה":"מילישניות"},decimal:"."},hr:{y:function(e){return e%10===2||e%10===3||e%10===4?"godine":"godina"},mo:function(e){return e===1?"mjesec":e===2||e===3||e===4?"mjeseca":"mjeseci"},w:function(e){return e%10===1&&e!==11?"tjedan":"tjedna"},d:function(e){return e===1?"dan":"dana"},h:function(e){return e===1?"sat":e===2||e===3||e===4?"sata":"sati"},m:function(e){var t=e%10;return(t===2||t===3||t===4)&&(e<10||e>14)?"minute":"minuta"},s:function(e){var t=e%10;return t===5||Math.floor(e)===e&&e>=10&&e<=19?"sekundi":t===1?"sekunda":t===2||t===3||t===4?"sekunde":"sekundi"},ms:function(e){return e===1?"milisekunda":e%10===2||e%10===3||e%10===4?"milisekunde":"milisekundi"},decimal:","},hi:{y:"साल",mo:function(e){return e===1?"महीना":"महीने"},w:function(e){return e===1?"हफ़्ता":"हफ्ते"},d:"दिन",h:function(e){return e===1?"घंटा":"घंटे"},m:"मिनट",s:"सेकंड",ms:"मिलीसेकंड",decimal:"."},hu:{y:"év",mo:"hónap",w:"hét",d:"nap",h:"óra",m:"perc",s:"másodperc",ms:"ezredmásodperc",decimal:","},id:{y:"tahun",mo:"bulan",w:"minggu",d:"hari",h:"jam",m:"menit",s:"detik",ms:"milidetik",decimal:"."},is:{y:"ár",mo:function(e){return"mánuð"+(e===1?"ur":"ir")},w:function(e){return"vik"+(e===1?"a":"ur")},d:function(e){return"dag"+(e===1?"ur":"ar")},h:function(e){return"klukkutím"+(e===1?"i":"ar")},m:function(e){return"mínút"+(e===1?"a":"ur")},s:function(e){return"sekúnd"+(e===1?"a":"ur")},ms:function(e){return"millisekúnd"+(e===1?"a":"ur")},decimal:"."},it:{y:function(e){return"ann"+(e===1?"o":"i")},mo:function(e){return"mes"+(e===1?"e":"i")},w:function(e){return"settiman"+(e===1?"a":"e")},d:function(e){return"giorn"+(e===1?"o":"i")},h:function(e){return"or"+(e===1?"a":"e")},m:function(e){return"minut"+(e===1?"o":"i")},s:function(e){return"second"+(e===1?"o":"i")},ms:function(e){return"millisecond"+(e===1?"o":"i")},decimal:","},ja:{y:"年",mo:"ヶ月",w:"週",d:"日",h:"時間",m:"分",s:"秒",ms:"ミリ秒",decimal:"."},km:{y:"ឆ្នាំ",mo:"ខែ",w:"សប្តាហ៍",d:"ថ្ងៃ",h:"ម៉ោង",m:"នាទី",s:"វិនាទី",ms:"មិល្លីវិនាទី"},kn:{y:function(e){return e===1?"ವರ್ಷ":"ವರ್ಷಗಳು"},mo:function(e){return e===1?"ತಿಂಗಳು":"ತಿಂಗಳುಗಳು"},w:function(e){return e===1?"ವಾರ":"ವಾರಗಳು"},d:function(e){return e===1?"ದಿನ":"ದಿನಗಳು"},h:function(e){return e===1?"ಗಂಟೆ":"ಗಂಟೆಗಳು"},m:function(e){return e===1?"ನಿಮಿಷ":"ನಿಮಿಷಗಳು"},s:function(e){return e===1?"ಸೆಕೆಂಡ್":"ಸೆಕೆಂಡುಗಳು"},ms:function(e){return e===1?"ಮಿಲಿಸೆಕೆಂಡ್":"ಮಿಲಿಸೆಕೆಂಡುಗಳು"}},ko:{y:"년",mo:"개월",w:"주일",d:"일",h:"시간",m:"분",s:"초",ms:"밀리 초",decimal:"."},ku:{y:"sal",mo:"meh",w:"hefte",d:"roj",h:"seet",m:"deqe",s:"saniye",ms:"mîlîçirk",decimal:","},lo:{y:"ປີ",mo:"ເດືອນ",w:"ອາທິດ",d:"ມື້",h:"ຊົ່ວໂມງ",m:"ນາທີ",s:"ວິນາທີ",ms:"ມິນລິວິນາທີ",decimal:","},lt:{y:function(e){return e%10===0||e%100>=10&&e%100<=20?"metų":"metai"},mo:function(e){return["mėnuo","mėnesiai","mėnesių"][C(e)]},w:function(e){return["savaitė","savaitės","savaičių"][C(e)]},d:function(e){return["diena","dienos","dienų"][C(e)]},h:function(e){return["valanda","valandos","valandų"][C(e)]},m:function(e){return["minutė","minutės","minučių"][C(e)]},s:function(e){return["sekundė","sekundės","sekundžių"][C(e)]},ms:function(e){return["milisekundė","milisekundės","milisekundžių"][C(e)]},decimal:","},lv:{y:function(e){return F(e)?"gads":"gadi"},mo:function(e){return F(e)?"mēnesis":"mēneši"},w:function(e){return F(e)?"nedēļa":"nedēļas"},d:function(e){return F(e)?"diena":"dienas"},h:function(e){return F(e)?"stunda":"stundas"},m:function(e){return F(e)?"minūte":"minūtes"},s:function(e){return F(e)?"sekunde":"sekundes"},ms:function(e){return F(e)?"milisekunde":"milisekundes"},decimal:","},mk:{y:function(e){return e===1?"година":"години"},mo:function(e){return e===1?"месец":"месеци"},w:function(e){return e===1?"недела":"недели"},d:function(e){return e===1?"ден":"дена"},h:function(e){return e===1?"час":"часа"},m:function(e){return e===1?"минута":"минути"},s:function(e){return e===1?"секунда":"секунди"},ms:function(e){return e===1?"милисекунда":"милисекунди"},decimal:","},mn:{y:"жил",mo:"сар",w:"долоо хоног",d:"өдөр",h:"цаг",m:"минут",s:"секунд",ms:"миллисекунд",decimal:"."},mr:{y:function(e){return e===1?"वर्ष":"वर्षे"},mo:function(e){return e===1?"महिना":"महिने"},w:function(e){return e===1?"आठवडा":"आठवडे"},d:"दिवस",h:"तास",m:function(e){return e===1?"मिनिट":"मिनिटे"},s:"सेकंद",ms:"मिलिसेकंद"},ms:{y:"tahun",mo:"bulan",w:"minggu",d:"hari",h:"jam",m:"minit",s:"saat",ms:"milisaat",decimal:"."},nl:{y:"jaar",mo:function(e){return e===1?"maand":"maanden"},w:function(e){return e===1?"week":"weken"},d:function(e){return e===1?"dag":"dagen"},h:"uur",m:function(e){return e===1?"minuut":"minuten"},s:function(e){return e===1?"seconde":"seconden"},ms:function(e){return e===1?"milliseconde":"milliseconden"},decimal:","},no:{y:"år",mo:function(e){return"måned"+(e===1?"":"er")},w:function(e){return"uke"+(e===1?"":"r")},d:function(e){return"dag"+(e===1?"":"er")},h:function(e){return"time"+(e===1?"":"r")},m:function(e){return"minutt"+(e===1?"":"er")},s:function(e){return"sekund"+(e===1?"":"er")},ms:function(e){return"millisekund"+(e===1?"":"er")},decimal:","},pl:{y:function(e){return["rok","roku","lata","lat"][x(e)]},mo:function(e){return["miesiąc","miesiąca","miesiące","miesięcy"][x(e)]},w:function(e){return["tydzień","tygodnia","tygodnie","tygodni"][x(e)]},d:function(e){return["dzień","dnia","dni","dni"][x(e)]},h:function(e){return["godzina","godziny","godziny","godzin"][x(e)]},m:function(e){return["minuta","minuty","minuty","minut"][x(e)]},s:function(e){return["sekunda","sekundy","sekundy","sekund"][x(e)]},ms:function(e){return["milisekunda","milisekundy","milisekundy","milisekund"][x(e)]},decimal:","},pt:{y:function(e){return"ano"+(e===1?"":"s")},mo:function(e){return e===1?"mês":"meses"},w:function(e){return"semana"+(e===1?"":"s")},d:function(e){return"dia"+(e===1?"":"s")},h:function(e){return"hora"+(e===1?"":"s")},m:function(e){return"minuto"+(e===1?"":"s")},s:function(e){return"segundo"+(e===1?"":"s")},ms:function(e){return"milissegundo"+(e===1?"":"s")},decimal:","},ro:{y:function(e){return e===1?"an":"ani"},mo:function(e){return e===1?"lună":"luni"},w:function(e){return e===1?"săptămână":"săptămâni"},d:function(e){return e===1?"zi":"zile"},h:function(e){return e===1?"oră":"ore"},m:function(e){return e===1?"minut":"minute"},s:function(e){return e===1?"secundă":"secunde"},ms:function(e){return e===1?"milisecundă":"milisecunde"},decimal:","},ru:{y:function(e){return["лет","год","года"][l(e)]},mo:function(e){return["месяцев","месяц","месяца"][l(e)]},w:function(e){return["недель","неделя","недели"][l(e)]},d:function(e){return["дней","день","дня"][l(e)]},h:function(e){return["часов","час","часа"][l(e)]},m:function(e){return["минут","минута","минуты"][l(e)]},s:function(e){return["секунд","секунда","секунды"][l(e)]},ms:function(e){return["миллисекунд","миллисекунда","миллисекунды"][l(e)]},decimal:","},sq:{y:function(e){return e===1?"vit":"vjet"},mo:"muaj",w:"javë",d:"ditë",h:"orë",m:function(e){return"minut"+(e===1?"ë":"a")},s:function(e){return"sekond"+(e===1?"ë":"a")},ms:function(e){return"milisekond"+(e===1?"ë":"a")},decimal:","},sr:{y:function(e){return["години","година","године"][l(e)]},mo:function(e){return["месеци","месец","месеца"][l(e)]},w:function(e){return["недељи","недеља","недеље"][l(e)]},d:function(e){return["дани","дан","дана"][l(e)]},h:function(e){return["сати","сат","сата"][l(e)]},m:function(e){return["минута","минут","минута"][l(e)]},s:function(e){return["секунди","секунда","секунде"][l(e)]},ms:function(e){return["милисекунди","милисекунда","милисекунде"][l(e)]},decimal:","},ta:{y:function(e){return e===1?"வருடம்":"ஆண்டுகள்"},mo:function(e){return e===1?"மாதம்":"மாதங்கள்"},w:function(e){return e===1?"வாரம்":"வாரங்கள்"},d:function(e){return e===1?"நாள்":"நாட்கள்"},h:function(e){return e===1?"மணி":"மணிநேரம்"},m:function(e){return"நிமிட"+(e===1?"ம்":"ங்கள்")},s:function(e){return"வினாடி"+(e===1?"":"கள்")},ms:function(e){return"மில்லி விநாடி"+(e===1?"":"கள்")}},te:{y:function(e){return"సంవత్స"+(e===1?"రం":"రాల")},mo:function(e){return"నెల"+(e===1?"":"ల")},w:function(e){return e===1?"వారం":"వారాలు"},d:function(e){return"రోజు"+(e===1?"":"లు")},h:function(e){return"గంట"+(e===1?"":"లు")},m:function(e){return e===1?"నిమిషం":"నిమిషాలు"},s:function(e){return e===1?"సెకను":"సెకన్లు"},ms:function(e){return e===1?"మిల్లీసెకన్":"మిల్లీసెకన్లు"}},uk:{y:function(e){return["років","рік","роки"][l(e)]},mo:function(e){return["місяців","місяць","місяці"][l(e)]},w:function(e){return["тижнів","тиждень","тижні"][l(e)]},d:function(e){return["днів","день","дні"][l(e)]},h:function(e){return["годин","година","години"][l(e)]},m:function(e){return["хвилин","хвилина","хвилини"][l(e)]},s:function(e){return["секунд","секунда","секунди"][l(e)]},ms:function(e){return["мілісекунд","мілісекунда","мілісекунди"][l(e)]},decimal:","},ur:{y:"سال",mo:function(e){return e===1?"مہینہ":"مہینے"},w:function(e){return e===1?"ہفتہ":"ہفتے"},d:"دن",h:function(e){return e===1?"گھنٹہ":"گھنٹے"},m:"منٹ",s:"سیکنڈ",ms:"ملی سیکنڈ",decimal:"."},sk:{y:function(e){return["rok","roky","roky","rokov"][g(e)]},mo:function(e){return["mesiac","mesiace","mesiace","mesiacov"][g(e)]},w:function(e){return["týždeň","týždne","týždne","týždňov"][g(e)]},d:function(e){return["deň","dni","dni","dní"][g(e)]},h:function(e){return["hodina","hodiny","hodiny","hodín"][g(e)]},m:function(e){return["minúta","minúty","minúty","minút"][g(e)]},s:function(e){return["sekunda","sekundy","sekundy","sekúnd"][g(e)]},ms:function(e){return["milisekunda","milisekundy","milisekundy","milisekúnd"][g(e)]},decimal:","},sl:{y:function(e){return e%10===1?"leto":e%100===2?"leti":e%100===3||e%100===4||Math.floor(e)!==e&&e%100<=5?"leta":"let"},mo:function(e){return e%10===1?"mesec":e%100===2||Math.floor(e)!==e&&e%100<=5?"meseca":e%10===3||e%10===4?"mesece":"mesecev"},w:function(e){return e%10===1?"teden":e%10===2||Math.floor(e)!==e&&e%100<=4?"tedna":e%10===3||e%10===4?"tedne":"tednov"},d:function(e){return e%100===1?"dan":"dni"},h:function(e){return e%10===1?"ura":e%100===2?"uri":e%10===3||e%10===4||Math.floor(e)!==e?"ure":"ur"},m:function(e){return e%10===1?"minuta":e%10===2?"minuti":e%10===3||e%10===4||Math.floor(e)!==e&&e%100<=4?"minute":"minut"},s:function(e){return e%10===1?"sekunda":e%100===2?"sekundi":e%100===3||e%100===4||Math.floor(e)!==e?"sekunde":"sekund"},ms:function(e){return e%10===1?"milisekunda":e%100===2?"milisekundi":e%100===3||e%100===4||Math.floor(e)!==e?"milisekunde":"milisekund"},decimal:","},sv:{y:"år",mo:function(e){return"månad"+(e===1?"":"er")},w:function(e){return"veck"+(e===1?"a":"or")},d:function(e){return"dag"+(e===1?"":"ar")},h:function(e){return"timm"+(e===1?"e":"ar")},m:function(e){return"minut"+(e===1?"":"er")},s:function(e){return"sekund"+(e===1?"":"er")},ms:function(e){return"millisekund"+(e===1?"":"er")},decimal:","},sw:{y:function(e){return e===1?"mwaka":"miaka"},mo:function(e){return e===1?"mwezi":"miezi"},w:"wiki",d:function(e){return e===1?"siku":"masiku"},h:function(e){return e===1?"saa":"masaa"},m:"dakika",s:"sekunde",ms:"milisekunde",decimal:".",_numberFirst:!0},tr:{y:"yıl",mo:"ay",w:"hafta",d:"gün",h:"saat",m:"dakika",s:"saniye",ms:"milisaniye",decimal:","},th:{y:"ปี",mo:"เดือน",w:"สัปดาห์",d:"วัน",h:"ชั่วโมง",m:"นาที",s:"วินาที",ms:"มิลลิวินาที",decimal:"."},vi:{y:"năm",mo:"tháng",w:"tuần",d:"ngày",h:"giờ",m:"phút",s:"giây",ms:"mili giây",decimal:","},zh_CN:{y:"年",mo:"个月",w:"周",d:"天",h:"小时",m:"分钟",s:"秒",ms:"毫秒",decimal:"."},zh_TW:{y:"年",mo:"個月",w:"周",d:"天",h:"小時",m:"分鐘",s:"秒",ms:"毫秒",decimal:"."}};function f(e){var t=function(a,m){var w=M({},t,m||{});return E(a,w)};return M(t,{language:"en",spacer:" ",conjunction:"",serialComma:!0,units:["y","mo","w","d","h","m","s"],languages:{},round:!1,unitMeasures:{y:315576e5,mo:26298e5,w:6048e5,d:864e5,h:36e5,m:6e4,s:1e3,ms:1}},e)}var p=f({});function y(e){var t=[e.language];if(s(e,"fallbacks"))if(T(e.fallbacks)&&e.fallbacks.length)t=t.concat(e.fallbacks);else throw new Error("fallbacks must be an array with at least one element");for(var i=0;i<t.length;i++){var a=t[i];if(s(e.languages,a))return e.languages[a];if(s(o,a))return o[a]}throw new Error("No language found.")}function E(e,t){var i,a,m;e=Math.abs(e);var w=y(t),k=[],b,S,R;for(i=0,a=t.units.length;i<a;i++){if(b=t.units[i],S=t.unitMeasures[b],i+1===a)if(s(t,"maxDecimalPoints")){var D=Math.pow(10,t.maxDecimalPoints),I=e/S;R=parseFloat((Math.floor(D*I)/D).toFixed(t.maxDecimalPoints))}else R=e/S;else R=Math.floor(e/S);k.push({unitCount:R,unitName:b}),e-=R*S}var P=0;for(i=0;i<k.length;i++)if(k[i].unitCount){P=i;break}if(t.round){var q,O;for(i=k.length-1;i>=0&&(m=k[i],m.unitCount=Math.round(m.unitCount),i!==0);i--)O=k[i-1],q=t.unitMeasures[O.unitName]/t.unitMeasures[m.unitName],(m.unitCount%q===0||t.largest&&t.largest-1<i-P)&&(O.unitCount+=m.unitCount/q,m.unitCount=0)}var L=[];for(i=0,k.length;i<a&&(m=k[i],m.unitCount&&L.push(v(m.unitCount,m.unitName,w,t)),L.length!==t.largest);i++);if(L.length){var z;if(s(t,"delimiter")?z=t.delimiter:s(w,"delimiter")?z=w.delimiter:z=", ",!t.conjunction||L.length===1)return L.join(z);if(L.length===2)return L.join(t.conjunction);if(L.length>2)return L.slice(0,-1).join(z)+(t.serialComma?",":"")+t.conjunction+L.slice(-1)}else return v(0,t.units[t.units.length-1],w,t)}function v(e,t,i,a){var m;s(a,"decimal")?m=a.decimal:s(i,"decimal")?m=i.decimal:m=".";var w;typeof i._formatCount=="function"?w=i._formatCount(e,m):w=e.toString().replace(".",m);var k=i[t],b;return typeof k=="function"?b=k(e):b=k,i._numberFirst?b+a.spacer+w:w+a.spacer+b}function M(e){for(var t,i=1;i<arguments.length;i++){t=arguments[i];for(var a in t)s(t,a)&&(e[a]=t[a])}return e}function h(e){return e===1?0:e===2?1:e>2&&e<11?2:0}function x(e){return e===1?0:Math.floor(e)!==e?1:e%10>=2&&e%10<=4&&!(e%100>10&&e%100<20)?2:3}function l(e){return Math.floor(e)!==e?2:e%100>=5&&e%100<=20||e%10>=5&&e%10<=9||e%10===0?0:e%10===1?1:e>1?2:0}function g(e){return e===1?0:Math.floor(e)!==e?1:e%10>=2&&e%10<=4&&e%100<10?2:3}function C(e){return e===1||e%10===1&&e%100>20?0:Math.floor(e)!==e||e%10>=2&&e%100>20||e%10>=2&&e%100<10?1:2}function F(e){return e%10===1&&e%100!==11}var T=Array.isArray||function(e){return Object.prototype.toString.call(e)==="[object Array]"};function s(e,t){return Object.prototype.hasOwnProperty.call(e,t)}p.getSupportedLanguages=function(){var t=[];for(var i in o)s(o,i)&&i!=="gr"&&t.push(i);return t},p.humanizer=f,u.exports?u.exports=p:this.humanizeDuration=p})()})(he);_.register();const U=document.querySelector("media-preview-modal"),c=class{constructor(r,n){d(this,"selectedFilesCardElement");d(this,"selectedFilesContainerElement");d(this,"selectedFilesRow");d(this,"previewBoxElement");d(this,"nameBoxElement");d(this,"sizeBoxElement");d(this,"typeBoxElement");d(this,"urlBoxElement");d(this,"statusBoxElement");d(this,"progressBoxElement");d(this,"progressBarElement");d(this,"progressSpeedBoxElement");d(this,"removeFileBoxElement");d(this,"fileTypeCategory");d(this,"progressAlertElement");d(this,"uploadStartTime",new Date().getTime());d(this,"file");d(this,"id");d(this,"_status");d(this,"_errorMessage","");d(this,"xhr");if(this.selectedFilesCardElement=r,this.file=n,this.id=c.nextFileID,this._status="pending",this.initializeSelectors(),this.populateSelectedFileRow(),this.addEventListeners(),c.uploadFilesMap.set(this.id,this),c.nextFileID++,!c.uploadButtonElement===null)throw new Error("upload button not found");c.uploadButtonVisible||(c.uploadButtonElement.classList.remove("opacity-0"),c.uploadButtonVisible=!0),c.uploadButtonDisabled&&(c.uploadButtonElement.disabled=!1,c.uploadButtonElement.classList.remove("disabled"),c.uploadButtonDisabled=!1),this.uploadAllPendingFilesListener()}get errorMessage(){return this._errorMessage}get status(){return this._status}addEventListeners(){this.addMediaPreviewClickListener(),this.addRemoveButtonListener()}initializeSelectors(){const r=document.getElementById("selected-files-template");if(r===null||r.content.firstElementChild===null)throw new Error("selected-files-template not found");this.selectedFilesRow=r.content.firstElementChild.cloneNode(!0),this.selectedFilesRow.dataset.fileId=this.id.toString(),this.previewBoxElement=this.selectedFilesRow.querySelector(".selected-file-preview-box"),this.nameBoxElement=this.selectedFilesRow.querySelector(".selected-file-name-box"),this.sizeBoxElement=this.selectedFilesRow.querySelector(".selected-file-size-box"),this.urlBoxElement=this.selectedFilesRow.querySelector(".selected-file-url-box"),this.statusBoxElement=this.selectedFilesRow.querySelector(".selected-file-status-box"),this.progressBoxElement=this.selectedFilesRow.querySelector(".selected-file-progress-box"),this.progressSpeedBoxElement=this.selectedFilesRow.querySelector(".selected-file-progress-speed-box"),this.removeFileBoxElement=this.selectedFilesRow.querySelector(".selected-file-remove-box"),this.typeBoxElement=this.selectedFilesRow.querySelector(".selected-file-type-box"),this.progressBarElement=this.progressBoxElement.querySelector("progress"),this.progressAlertElement=this.progressBoxElement.querySelector("fileshare-alert"),this.selectedFilesContainerElement=document.getElementById("selected-files-container");const n=[];if(this.previewBoxElement===null&&n.push("selected-file-preview-box"),this.nameBoxElement===null&&n.push("selected-file-name-box"),this.sizeBoxElement===null&&n.push("selected-file-size-box"),this.urlBoxElement===null&&n.push("selected-file-url-box"),this.statusBoxElement===null&&n.push("selected-file-status-box"),this.progressBoxElement===null&&n.push("selected-file-progress-box"),this.progressSpeedBoxElement===null&&n.push("selected-file-progress-speed-box"),this.removeFileBoxElement===null&&n.push("selected-file-remove-box"),this.progressBarElement===null&&n.push("selected-file-progress-box progress"),this.selectedFilesContainerElement===null&&n.push("selected-files-container"),n.length>0)throw new Error("template elements not found: "+n.join(", "))}populateSelectedFileRow(){switch(this.nameBoxElement.textContent=this.file.name,this.statusBoxElement.textContent=this._status,this.typeBoxElement.textContent=this.file.type,this.sizeBoxElement.textContent=Y(this.file.size,{output:"string"}),this.file.type.startsWith("image/")?this.fileTypeCategory="image":this.file.type.startsWith("video/")?this.fileTypeCategory="video":this.file.type.startsWith("audio/")?this.fileTypeCategory="audio":this.fileTypeCategory="other",this.fileTypeCategory){case"image":const r=new FileReader;r.onload=p=>{const y=p.target;if(y==null)throw new Error("file target reader is null");const E=y.result;if(E==null)throw new Error("file dataURL is null");const v=document.createElement("img");v.src=E,this.previewBoxElement.appendChild(v)},r.readAsDataURL(this.file);break;case"video":const n=document.createElement("video");n.src=URL.createObjectURL(this.file),n.controls=!1,this.previewBoxElement.appendChild(n);break;case"audio":const o=document.createElement("audio");o.src=URL.createObjectURL(this.file),o.controls=!0,this.previewBoxElement.appendChild(o);break;default:const f=document.createElement("img");f.src="/static/images/no_preview.webp",this.previewBoxElement.appendChild(f)}this.selectedFilesContainerElement.prepend(this.selectedFilesRow),this.selectedFilesCardElement.classList.remove("hidden"),this.selectedFilesRow.offsetWidth}addRemoveButtonListener(){this.removeFileBoxElement.addEventListener("click",()=>{this.selectedFilesRow.style.animationName="fadeOutRightBig",this.selectedFilesRow.style.animationDuration="0.4s",setTimeout(()=>{c.uploadFilesMap.delete(this.id),this.selectedFilesRow.remove(),c.uploadFilesMap.size===0&&(c.uploadButtonElement.classList.remove("disabled"),c.uploadButtonElement.classList.add("opacity-0"),c.uploadButtonVisible=!1,this.selectedFilesCardElement.classList.add("hidden"))},350)})}uploadAllPendingFilesListener(){c.uploadButtonElement.addEventListener("click",()=>{c.uploadButtonVisible&&(c.uploadButtonDisabled||c.uploadAllSelectedFiles())})}addMediaPreviewClickListener(){this.previewBoxElement.addEventListener("click",()=>{var n;const r=(n=U.shadowRoot)==null?void 0:n.querySelector("#media-preview-slot");switch(this.fileTypeCategory){case"video":const o=this.previewBoxElement.querySelector("video");if(o===null)return;const f=o.cloneNode(!0);f.controls=!0,r.replaceChildren(f),U.showModal();break;case"audio":const p=this.previewBoxElement.querySelector("audio");if(p===null)return;const y=p.cloneNode(!0);y.controls=!0,r.replaceChildren(y),U.showModal();break;default:const E=this.previewBoxElement.querySelector("img");if(E===null)return;const v=E.cloneNode(!0);r.replaceChildren(v),U.showModal();break}})}calculateSpeedBetweenProgressUpdates(r,n,o,f){const p=n-r,y=f-o;return this.calculateSpeedPerSecond(p,y)}calculateSpeedPerSecond(r,n){return r/n}uploadFilePromise(){return this.setStatus("uploading"),this.progressBarElement.classList.remove("hidden"),this.progressSpeedBoxElement.classList.remove("hidden"),this.progressAlertElement.classList.add("hidden"),new Promise((r,n)=>{let o=0,f=new Date().getTime();this.uploadStartTime=new Date().getTime(),this.xhr=new XMLHttpRequest,this.xhr.open("POST",c.uploadURL,!0),this.xhr.responseType="json",this.xhr.onload=()=>{this.xhr.status===200?r(null):n()},this.xhr.onerror=()=>{n()},this.xhr.upload.onprogress=y=>{if(y.lengthComputable){this.progressBarElement.value=y.loaded,this.progressBarElement.max=y.total;const E=Math.round(f/1e3),v=Math.round(new Date().getTime()/1e3),M=this.calculateSpeedBetweenProgressUpdates(o,y.loaded,E,v);this.progressSpeedBoxElement.textContent=`${Y(M,{output:"string"})}/s`}},this.xhr.upload.onloadend=y=>{const E=new Date().getTime()-this.uploadStartTime,v=V(E,{round:!0}),M=this.calculateSpeedPerSecond(y.loaded,E/1e3),h=`${Y(M,{output:"string"})}/s`;this.progressSpeedBoxElement.textContent=`Took ${v} @ ${h}`};const p=new FormData;p.append("file",this.file),this.xhr.send(p)})}handleUploadSuccess(){this.setStatus("done");const r=this.xhr.response;if(r.error){this.setStatus("error"),this.handleUploadFailure();return}const n=document.createElement("a");n.href=r.url,n.textContent=r.url,n.target="_blank",this.urlBoxElement.replaceChildren(n),this.progressBarElement.classList.add("hidden"),this.progressAlertElement.classList.remove("hidden"),this.progressAlertElement.classList.remove("alert-error"),this.progressAlertElement.classList.add("alert-success"),this.progressAlertElement.textContent="Upload successful"}handleUploadFailure(){this.setStatus("error"),this.progressBarElement.classList.add("hidden"),this.progressAlertElement.classList.remove("hidden"),this.progressAlertElement.classList.remove("alert-success"),this.progressAlertElement.classList.add("alert-error"),this.progressAlertElement.textContent="Failed to upload file! "+(this.xhr.responseText!==""?this.xhr.responseText:this.xhr.statusText),this._errorMessage=this.xhr.responseText!==""?this.xhr.responseText:this.xhr.statusText}setStatus(r){switch(this._status=r,this._status){case"pending":this.statusBoxElement.textContent="Pending";break;case"uploading":this.statusBoxElement.textContent="Uploading";break;case"done":this.statusBoxElement.textContent="Completed";break;case"error":this.statusBoxElement.textContent="Error"}}static uploadAllSelectedFiles(){c.uploadButtonElement.disabled=!0,c.uploadButtonElement.classList.add("disabled"),c.uploadButtonDisabled=!0;let r=Array();if(c.uploadFilesMap.forEach(n=>{n._status==="pending"&&r.push(n)}),r.length!==0)for(let n=0;n<c.simultaneousUploads;n++){const o=async()=>{const f=r.shift();if(f)try{await f.uploadFilePromise(),f.handleUploadSuccess()}catch{f.handleUploadFailure()}finally{await o()}};o()}}};let B=c;d(B,"nextFileID",0),d(B,"uploadFilesMap",new Map),d(B,"simultaneousUploads",3),d(B,"uploadURL","/upload"),d(B,"uploadButtonElement",document.getElementById("selected-files-upload-button")),d(B,"uploadButtonVisible",!1),d(B,"uploadButtonDisabled",!1);function pe(u){const r=document.createDocumentFragment(),n=document.createElement("div");n.classList.add("full-screen-loader");const o=document.createElement("i");o.className="fa-solid fa-circle-notch fa-spin";const f=document.createElement("span");f.textContent=u,n.appendChild(o),n.appendChild(f),r.appendChild(n),document.documentElement.appendChild(r),console.log(r)}function J(){const u=document.querySelector(".full-screen-loader");u&&u.remove()}const ge=document.getElementById("upload-zone-input"),$=document.getElementById("selected-files-card");document.documentElement.addEventListener("drag",u=>{u.preventDefault()});document.documentElement.addEventListener("dragenter",u=>{document.documentElement.classList.add("drag-drop-action"),u.preventDefault()});document.documentElement.addEventListener("dragleave",u=>{document.documentElement.classList.remove("drag-drop-action"),u.preventDefault()});document.documentElement.addEventListener("dragover",u=>{document.documentElement.classList.add("drag-drop-action"),u.preventDefault()});document.documentElement.addEventListener("drop",u=>{document.documentElement.classList.remove("drag-drop-action"),u.preventDefault();const r=u.dataTransfer;r!=null&&Array.from(r.files).forEach(n=>{new B($,n)})});ge.addEventListener("change",u=>{u.preventDefault();const n=u.target.files;Array.from(n).forEach(o=>{new B($,o)})});document.addEventListener("paste",u=>{pe("Pasting files...");const r=u.clipboardData;if(r==null){J();return}u.preventDefault(),Array.from(r.files).forEach(n=>{new B($,n)}),J()});