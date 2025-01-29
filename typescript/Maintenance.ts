export type MaintenanceType =
  | 'Plumbing'
  | 'Electrical'
  | 'General'
  | 'Gardening'
  | 'Pool'
  | 'Painting'
  | 'Exterminator';

export interface Maintenance {
  id: string; // UUID
  unitGroupId?: string; // UUID
  unitId?: string; // UUID
  type: MaintenanceType;
  comment: string;
  startDate: Date;
  endDate: Date;
  isDone: boolean;
  createdAt: Date;
  updatedAt: Date;
}
