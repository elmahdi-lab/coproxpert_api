export type Role = 1 | 2 | 3;
export type EntityName = 'unit_group' | 'unit';

export interface Permission {
  id: string; // UUID
  userId: string; // UUID
  entityId: string; // UUID
  entityName: EntityName;
  role: Role;
  createdAt: Date;
  updatedAt: Date;
}
