#include "termbox2/termbox2.h"


int main() {
    tb_init(); // Inicjalizacja Termbox

    // Ustaw kolor tekstu na kolor z palety indeksowanych (np. niebieski)
    tb_set_cell(10, 10, 'A', TB_WHITE, TB_BLUE);

    // Wyświetlenie okna na ekranie
    tb_present();

    // Oczekiwanie na klawisz
    tb_poll_event(NULL);

    // Zamknięcie Termbox
    tb_shutdown();

    return 0;
}
