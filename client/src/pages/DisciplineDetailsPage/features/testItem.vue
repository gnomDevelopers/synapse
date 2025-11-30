<template>
  <div class="bg-white border border-gray-200 rounded-lg p-4 flex flex-col sm:flex-row sm:items-center gap-4 hover:border-gray-400 transition-colors">
    <div class="w-10 h-10 bg-gray-100 rounded-full flex items-center justify-center shrink-0 text-gray-500">
      <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
    </div>

    <div class="grow cursor-pointer">
      <h3 class="text-gray-800 font-medium select-none">{{ title }}</h3>
      <div class="text-sm text-gray-500 mt-1 flex flex-wrap gap-x-4">
        <span class="select-none">Дедлайн: <span class="text-gray-700 font-semibold">{{ deadline }}</span></span>
      </div>
    </div>

    <div class="shrink-0 flex items-center gap-2 bg-gray-50 px-3 py-1 rounded border border-gray-200 select-none">
      <span class="text-sm text-gray-500 font-bold">Проходной:</span>
      <span class="text-gray-800 font-medium">{{ passingScore }}/{{ maxScore }}</span>
    </div>

    <div class="flex flex-row items-center gap-2">
      <button class="text-gray-400 hover:text-gray-600 p-2 cursor-pointer select-none">
        <svg class="w-5.5 h-5.5" fill="none" stroke="currentColor" viewBox="0 0 35 27" xmlns="http://www.w3.org/2000/svg">
          <path d="M1.4585 13.1667C1.4585 13.1667 7.29183 1.5 17.5002 1.5C27.7085 1.5 33.5418 13.1667 33.5418 13.1667C33.5418 13.1667 27.7085 24.8333 17.5002 24.8333C7.29183 24.8333 1.4585 13.1667 1.4585 13.1667Z" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
          <path d="M17.5002 17.5417C19.9164 17.5417 21.8752 15.5829 21.8752 13.1667C21.8752 10.7504 19.9164 8.79167 17.5002 8.79167C15.0839 8.79167 13.1252 10.7504 13.1252 13.1667C13.1252 15.5829 15.0839 17.5417 17.5002 17.5417Z" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
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
    deadline: {
      type: String,
      default: '30 ноября 2025 г.',
    },
    passingScore: {
      type: Number,
      default: 0,
    },
    maxScore: {
      type: Number,
      default: 100,
    },
  }
}
</script>