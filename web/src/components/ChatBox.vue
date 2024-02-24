<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import ChatInput from './ChatInput.vue';
import ChatMessages from './ChatMessages.vue'

interface ChatMessage {
    id: number;
    message: string;
    time: string;
    direction: number; // inbound 1, outbound 2
}

let websocket: WebSocket;
const messages = ref<ChatMessage[]>([]);

const initChat = () => {
    const chatStreamUrl = "wss://echo.websocket.org";
    websocket = new WebSocket(chatStreamUrl);

    websocket.onopen = () => {
        console.log('Connected to chat stream');
    }
    websocket.onmessage = (event) => {
        console.log('Received message:', event.data);
        const message = event.data;
        messages.value.push({
            id: messages.value.length + 1,
            message: message,
            time: new Date().toLocaleTimeString(),
            direction: 1
        });
    }
}

let onMessageSubmit = (message: string) => {
    console.log('Sending message:', message);
    const outboundMessage: ChatMessage = {
        id: messages.value.length + 1,
        message: message,
        time: new Date().toLocaleTimeString(),
        direction: 2
    }
    messages.value.push(outboundMessage);
}

onMounted(() => {
    initChat();
});

</script>
<template>

	<!-- Component Start -->
	<div class="flex flex-col flex-grow w-full max-w-xl bg-white shadow-xl rounded-lg overflow-hidden">
		
		<ChatMessages 
            :messages="messages"
        />
		<ChatInput 
            @submit="onMessageSubmit"
        />
	</div>
	<!-- Component End  -->
</template>

<style>
#chat-messages {
    height: 400px;
}
</style>
