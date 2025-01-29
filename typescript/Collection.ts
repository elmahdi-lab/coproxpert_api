export interface Collection {
  id: string; // UUID
  description: string;
  ownerId: string; // UUID
  unitId: string; // UUID
  amount: number;
  createdAt: Date;
  updatedAt: Date;
}
