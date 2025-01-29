export interface Budget {
  id: string; // UUID
  ownerId: string; // UUID
  unitGroupId: string; // UUID
  amount: number;
  year: number;
  isProvisional: boolean;
  createdAt: Date;
  updatedAt: Date;
}
