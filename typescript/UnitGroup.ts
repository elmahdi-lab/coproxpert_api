export interface UnitGroup {
  id: string; // UUID
  name: string;
  ownerId: string; // UUID
  owner?: User;
  createdAt: Date;
  updatedAt: Date;
}
