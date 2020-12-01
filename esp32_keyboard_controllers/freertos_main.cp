#include <stdio.h>
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"
//#include "driver/gpio.h"
#include "sdkconfig.h"
#include "esp_system.h"
#include "esp_spi_flash.h"

/* Can run 'make menuconfig' to choose the GPIO to blink,
   or you can edit the following line and set a number here.
*/

// Avoids the "undefined reference to app_main"
extern "C" {
void app_main(void);
}

//void blink_task(void *pvParameter)
//{
//        vTaskDelay(1000 / portTICK_PERIOD_MS);
//    }
//}

void app_main(void)
{
    //xTaskCreate(&blink_task, "blink_task", configMINIMAL_STACK_SIZE, NULL, 5, NULL);
    gpio_set_direction(GPIO_NUM_23, GPIO_MODE_OUTPUT);
    printf("Hello world \n");
    //while(1) {
    //    gpio_set_level(GPIO_NUM_23, 0);
    //    vTaskDelay(1000 / portTICK_PERIOD_MS);
    //    gpio_set_level(GPIO_NUM_23, 1);
    //    vTaskDelay(1000 / portTICK_PERIOD_MS);
    //}

        /* Print chip information */
    esp_chip_info_t chip_info;
    esp_chip_info(&chip_info);
    printf("This is ESP32 chip with %d CPU cores, WiFi%s%s, ",
            chip_info.cores,
            (chip_info.features & CHIP_FEATURE_BT) ? "/BT" : "",
            (chip_info.features & CHIP_FEATURE_BLE) ? "/BLE" : "");

    printf("silicon revision %d, ", chip_info.revision);

    printf("%dMB %s flash\n", spi_flash_get_chip_size() / (1024 * 1024),
            (chip_info.features & CHIP_FEATURE_EMB_FLASH) ? "embedded" : "external");

    for (int i = 10; i >= 0; i--) {
        printf("Restarting in %d seconds...\n", i);
        vTaskDelay(1000 / portTICK_PERIOD_MS);
    }

    printf("Restarting now.\n");
    fflush(stdout);
    esp_restart();

}