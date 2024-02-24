<script setup lang="ts">

import { ref } from 'vue';
import ChatMessage from './ChatMessage.vue';
const props = defineProps(['messages']);

</script>
<template>
    <div id="chat-messages" 
    class="flex flex-col flex-grow h-0 p-4 overflow-auto">
        <ChatMessage
            v-for="message in props.messages"
            :key="message.id"
            :message="message"
        />
    </div>
</template>