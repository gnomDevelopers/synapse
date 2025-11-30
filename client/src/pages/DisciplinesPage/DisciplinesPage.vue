<template>
  <div>
    <main class="grow flex flex-col items-center py-8 px-4">
      
      <div class="w-full max-w-4xl flex justify-between items-center mb-6">
        <h1 class="text-2xl font-bold text-gray-800 cursor-default">Мои дисциплины</h1>
        <iconButton class="btn-main gap-2" @click="openModal">
          <template #text>
            <p>Добавить</p>
          </template>
        </iconButton>
      </div>

      <div class="w-full flex flex-col items-center gap-6">
        <div v-if="isLoading" class="text-gray-500 select-none cursor-default">Загрузка...</div>
        
        <DisciplineItem 
          v-for="item in disciplines" 
          :key="item.id" 
          :discipline="item"
          @navigate="handleNavigate"
        />

        <div v-if="!isLoading && disciplines.length === 0" class="text-gray-400 mt-10">
          Список дисциплин пуст
        </div>
      </div>
    </main>

    <div v-if="isModalOpen" class="fixed inset-0 z-50 flex items-center justify-center">
      
      <div 
        class="absolute inset-0 bg-blur/20 bg-opacity-40 backdrop-blur-xs transition-opacity"
        @click="closeModal"
      ></div>

      <div class="relative bg-white rounded-lg shadow-xl w-full max-w-lg mx-4 p-6 z-10">
        <h2 class="text-xl font-bold text-gray-800 mb-4">Создать дисциплину</h2>
        
        <form @submit.prevent="createDiscipline" class="flex flex-col gap-4">
          
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Название</label>
            <input 
              v-model="newDiscipline.title"
              type="text" 
              required
              class="w-full border border-gray-300 rounded p-2 focus:ring-2 focus:ring-gray-400 focus:outline-none"
              placeholder="Введите название предмета"
            />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Описание</label>
            <textarea 
              v-model="newDiscipline.description"
              rows="3"
              required
              class="w-full border border-gray-300 rounded p-2 focus:ring-2 focus:ring-gray-400 focus:outline-none"
              placeholder="Краткое описание"
            ></textarea>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1">Обложка</label>
            <input 
              type="file" 
              accept="image/*"
              @change="handleFileChange"
              class="block w-full text-sm text-gray-500
                file:mr-4 file:py-2 file:px-4
                file:rounded file:border-0
                file:text-sm file:font-semibold
                file:bg-gray-100 file:text-gray-700
                hover:file:bg-gray-200"
            />
            <div v-if="previewImage" class="mt-2 w-full h-32 bg-gray-100 rounded overflow-hidden">
              <img :src="previewImage" class="w-full h-full object-cover" />
            </div>
          </div>

          <div class="flex justify-end gap-3 mt-4">
            <button 
              type="button" 
              @click="closeModal" 
              class="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
            >
              Отмена
            </button>
            <button 
              type="submit" 
              class="px-6 py-2 bg-gray-700 text-white rounded hover:bg-gray-800 transition-colors"
              :disabled="isSubmitting"
            >
              {{ isSubmitting ? 'Создание...' : 'Создать' }}
            </button>
          </div>

        </form>
      </div>
    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import DisciplineItem from './features/disciplineItem.vue';
import type { IDiscipline } from '@/utils/types';

interface INewDisciplineForm {
  title: string;
  description: string;
  imageFile: File | null;
}

export default defineComponent({
  name: 'DisciplinePage',
  components: {
    DisciplineItem
  },
  data() {
    return {
      // Список дисциплин
      disciplines: [] as IDiscipline[],
      
      // Состояние UI
      isLoading: false,
      isModalOpen: false,
      isSubmitting: false,

      // Данные формы
      newDiscipline: {
        title: '',
        description: '',
        imageFile: null
      } as INewDisciplineForm,
      
      previewImage: '' as string
    };
  },
  mounted() {
    this.fetchDisciplines();
    this.fetchUser();
  },
  methods: {
    // --- Методы API ---

    async fetchUser() {
      // API: GET /user/me
      // this.user = response.data
    },

    async fetchDisciplines() {
      this.isLoading = true;
      try {
        // API: GET /disciplines
        // const response = await api.get('/disciplines');
        // this.disciplines = response.data;
        
        // --- Временная имитация данных ---
        setTimeout(() => {
          this.disciplines = [
            { id: 1, title: 'Высшая математика', description: 'Курс линейной алгебры и аналитической геометрии.', imageUrl: null },
            { id: 2, title: 'Физика', description: 'Общий курс физики. Механика и термодинамика.', imageUrl: null },
          ];
          this.isLoading = false;
        }, 500);
        // --------------------------------
      } catch (error) {
        console.error(error);
        this.isLoading = false;
      }
    },

    async createDiscipline() {
      if (!this.newDiscipline.title || !this.newDiscipline.description) return;

      this.isSubmitting = true;
      
      try {
        // Подготовка данных для отправки (обычно FormData для файлов)
        const formData = new FormData();
        formData.append('title', this.newDiscipline.title);
        formData.append('description', this.newDiscipline.description);
        if (this.newDiscipline.imageFile) {
          formData.append('image', this.newDiscipline.imageFile);
        }

        // API: POST /disciplines
        // const response = await api.post('/disciplines', formData);
        // Добавляем созданную дисциплину в начало списка
        // this.disciplines.unshift(response.data);

        // --- Временная локальная логика ---
        const fakeNewDiscipline: IDiscipline = {
          id: Date.now(),
          title: this.newDiscipline.title,
          description: this.newDiscipline.description,
          imageUrl: this.previewImage || null // Используем превью как картинку для теста
        };
        this.disciplines.unshift(fakeNewDiscipline);
        // ---------------------------------

        this.closeModal();
      } catch (error) {
        console.error('Ошибка создания:', error);
      } finally {
        this.isSubmitting = false;
      }
    },

    // --- Методы UI и логики ---

    openModal() {
      this.isModalOpen = true;
    },

    closeModal() {
      this.isModalOpen = false;
      this.resetForm();
    },

    resetForm() {
      this.newDiscipline = {
        title: '',
        description: '',
        imageFile: null
      };
      this.previewImage = '';
    },

    handleFileChange(event: Event) {
      const input = event.target as HTMLInputElement;
      if (input.files && input.files[0]) {
        const file = input.files[0];
        this.newDiscipline.imageFile = file;
        
        // Создаем URL для превью
        this.previewImage = URL.createObjectURL(file);
      }
    },

    handleNavigate(id: number) {
      console.log(`Переход к дисциплине с ID: ${id}`);
      // router.push({ name: 'DisciplineDetails', params: { id } });
    }
  }
});
</script>