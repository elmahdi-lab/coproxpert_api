export interface Inspection {
  id: string; // UUID
  unitGroupId: string; // UUID
  ownerId: string; // UUID
  assignedTo: string; // UUID
  startDate: string; // Date
  endDate: string; // Date
  isComplete: boolean;
  isCompleteAt: Date;
  details: string;
  createdAt: Date;
  updatedAt: Date;
}
