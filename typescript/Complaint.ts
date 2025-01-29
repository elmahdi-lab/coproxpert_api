export interface Complaint {
  id: string; // UUID
  title: string;
  description: string;
  reporterId: string; // UUID
  unitGroupId: string; // UUID
  files: File[];
  isResolved: boolean;
  resolvedAt: string;
  response: string;
  createdAt: Date;
  updatedAt: Date;
}
