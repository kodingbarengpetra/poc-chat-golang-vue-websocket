<script setup lang="ts">
import { ref } from 'vue'

const chatMessage = ref('');
const emit = defineEmits(['submit'])

const onMessageSubmit = (message: string) => {
    emit('submit', message);
    chatMessage.value = '';
}

</script>
<template>
    <div class="bg-gray-300 p-4">
		<input 
            v-model="chatMessage"
            @keyup.enter="onMessageSubmit(chatMessage)"
            class="flex items-center h-10 w-full rounded px-3 text-sm" 
            type="text" 
            placeholder="Type your message…"
        >
    </div>
</template>