
// Интерфейс дисциплины
export interface IDiscipline {
  id: number;
  title: string;
  description: string;
  imageUrl: string | null;
}

// --- Интерфейсы ---
export interface IMaterial {
  id: number;
  title: string;
  fileUrl: string;
  createdAt: string; // ISO Date string
}

export interface IAssigment {
  id: number;
  title: string;
  fileUrl: string;
  createdAt: string; // ISO Date string
}

export interface ITest {
  id: number;
  title: string;
  deadline: string; // ISO Date string
  passingScore: number;
  maxScore: number;
}