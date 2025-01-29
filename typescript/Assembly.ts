export interface Assembly {
  id: string; // UUID
  name: string;
  ownerId: string; // UUID
  unitGroupId: string; // UUID
  startDate: string; // Date
  endDate: string; // Date
  active?: boolean;
}
