# seeduino xiao lamp
Лампа "встречающий свет" на микроконтроллере xiao samd21
* Включение при наступлении темноты (задается подстроечным резистором)
* Выключение после истечении заданного времени (задается подстроечным резистором)
* Контакты "сброс" (можно установить кнопку или другой прерыватель) - если лампа в данный момент выключена и темно, то включает освещение повторно; если лампа включена, то выключит её 
* Лампу нужно использовать до 0.2A, 5V. Токоограничивающий резистор не входит в схему.

xiao - https://wiki.seeedstudio.com/Seeeduino-XIAO/ \
tinygo - https://tinygo.org/docs/reference/microcontrollers/machine/xiao/

---
#### Структура: ####
 * https://github.com/LotauRus/xiao_lamp - TinyGo проект (tinygo version 0.30.0 windows/amd64 (using go version go1.21.1 and LLVM version 16.0.1)) 
 * https://github.com/LotauRus/xiao_lamp/bin - Готовая прошивка для xiao  
 * https://github.com/LotauRus/xiao_lamp/resources/kicad - Файлы проекта KiCAD (компоненты, использованные в схеме kicad, можно найти тут: https://github.com/LotauRus/kicad_3rdparty) 
 * https://github.com/LotauRus/xiao_lamp/resources/kicad/documents - Вспомогательная документация, даташиты 
 * https://github.com/LotauRus/xiao_lamp/resources/kicad/documents/pcb - Файлы для печати: проект, схема на плёнку под ЛУТ (фоторезист), схема сверления 
---
#### Компоненты: ####
 * Стабилизатор LM78M05_TO252
 * Резистор 1206 10K
 * Резистор 1206 100K
 * Конденсатор 1206 0.1мкф
 * Конденсатор 1206 0.33мкф
 * Транзистор SOT-23 2N7002
 * Коннектор JST_XH 1x02 P2.50mm_Vertical (4шт)
 * Микроконтроллер Seeduino XIAO samd21
 * Потенциометр Bourns_3296W_Vertical (2шт)