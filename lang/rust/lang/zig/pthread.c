#include <pthread.h>
#include <stdio.h>

void* print_message(void* message) {
    printf("%s\n", (char*)message);
    return NULL;
}

int main() {
    pthread_t thread;
    const char* message = "Hello from pthread!";

    // 创建线程
    if (pthread_create(&thread, NULL, print_message, (void*)message)) {
        fprintf(stderr, "Error creating thread\n");
        return 1;
    }

    // 等待线程结束
    pthread_join(thread, NULL);

    return 0;
}
