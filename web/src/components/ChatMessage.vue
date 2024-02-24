<script setup lang="ts">
const props = defineProps(['message']);
</script>
<template>
    <div class="flex w-full mt-2 space-x-3 max-w-xs" :class="props.message.direction == 1 ? '' : 'ml-auto justify-end'">
        <div class="flex-shrink-0 h-10 w-10 rounded-full bg-gray-300" v-if="props.message.direction == 1"></div>
        <div>
            <div class="p-3 rounded-r-lg rounded-bl-lg"
                :class="props.message.direction == 1 ? 'bg-gray-300 ' : 'bg-blue-600 text-white'"
            >
                <p class="text-sm">
                    {{ props.message.message }}
                </p>
            </div>
            <span class="text-xs text-gray-500 leading-none">
                {{ props.message.time }}
            </span>
        </div>

        <div class="flex-shrink-0 h-10 w-10 rounded-full bg-gray-300" v-if="props.message.direction == 2"></div>
    </div>
</template>

