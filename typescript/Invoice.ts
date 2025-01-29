export interface Invoice {
  id: string; // UUID
  amount: number;
  description: string;
  createdAt: Date;
  updatedAt: Date;
}
