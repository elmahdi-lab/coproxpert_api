export interface Fundraising {
  id: string; // UUID
  amount: number;
  description: string;
  ownerId: string; // UUID
  unitGroupId: string; // UUID
  createdAt: Date;
  updatedAt: Date;
}
