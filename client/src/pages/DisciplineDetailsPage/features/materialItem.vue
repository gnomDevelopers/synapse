<template>
  <div class="bg-white border border-gray-200 rounded-lg p-4 flex items-center gap-4 hover:shadow-md transition-shadow cursor-pointer">
    <div class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center shrink-0 text-gray-500">
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 27 33" xmlns="http://www.w3.org/2000/svg"><path d="M1.5 27.0208C1.5 26.0539 1.88411 25.1266 2.56784 24.4428C3.25157 23.7591 4.1789 23.375 5.14583 23.375H24.8333M1.5 27.0208C1.5 27.9878 1.88411 28.9151 2.56784 29.5988C3.25157 30.2826 4.1789 30.6667 5.14583 30.6667H24.8333V1.5H5.14583C4.1789 1.5 3.25157 1.88411 2.56784 2.56784C1.88411 3.25157 1.5 4.1789 1.5 5.14583V27.0208Z" stroke="#888383" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></svg>
    </div>
    
    <div class="grow">
      <h3 class="text-gray-800 font-medium">{{ title }}</h3>
      <span class="text-xs text-gray-500">Добавлено: {{ createdAt }}</span>
    </div>

    <div class="flex flex-row gap-2">
      <button class="text-gray-400 hover:text-gray-600 p-2 cursor-pointer select-none">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 30 30" xmlns="http://www.w3.org/2000/svg">
          <path d="M27.75 19V24.8333C27.75 25.6069 27.4427 26.3487 26.8957 26.8957C26.3487 27.4427 25.6069 27.75 24.8333 27.75H4.41667C3.64312 27.75 2.90125 27.4427 2.35427 26.8957C1.80729 26.3487 1.5 25.6069 1.5 24.8333V19M7.33333 11.7083L14.625 19M14.625 19L21.9167 11.7083M14.625 19V1.5" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>    
      </button>

      <button class="text-gray-400 hover:text-gray-600 p-2 cursor-pointer select-none">
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 33 33" xmlns="http://www.w3.org/2000/svg">
          <path d="M14.625 4.59359H4.41667C3.64312 4.59359 2.90125 4.90088 2.35427 5.44786C1.80729 5.99484 1.5 6.73671 1.5 7.51026V27.9269C1.5 28.7005 1.80729 29.4423 2.35427 29.9893C2.90125 30.5363 3.64312 30.8436 4.41667 30.8436H24.8333C25.6069 30.8436 26.3487 30.5363 26.8957 29.9893C27.4427 29.4423 27.75 28.7005 27.75 27.9269V17.7186M25.5625 2.40609C26.1427 1.82593 26.9295 1.5 27.75 1.5C28.5705 1.5 29.3573 1.82593 29.9375 2.40609C30.5177 2.98625 30.8436 3.77312 30.8436 4.59359C30.8436 5.41406 30.5177 6.20093 29.9375 6.78109L16.0833 20.6353L10.25 22.0936L11.7083 16.2603L25.5625 2.40609Z" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </button>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  props: {
    id: {
      type: Number,
      default: 0,
    },
    title: {
      type: String,
      default: '',
    },
    createdAt: {
      type: String,
      default: '30 ноября 2025 г.',
    },
  }
}
</script>