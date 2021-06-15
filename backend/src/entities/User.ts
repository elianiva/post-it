import {
  BaseEntity,
  BeforeInsert,
  BeforeUpdate,
  Column,
  Entity,
  PrimaryColumn,
} from "typeorm";

@Entity()
export class User extends BaseEntity {
  @PrimaryColumn({ nullable: false })
  id!: string;

  @Column({ nullable: false })
  email!: string;

  @Column({ nullable: false })
  username!: string;

  @Column({ nullable: false })
  password!: string;

  @Column({ nullable: false, name: "created_at", default: () => "NOW()" })
  createdAt!: Date;

  @Column({ nullable: false, name: "updated_at", default: () => "NOW()" })
  updatedAt!: Date;

  @BeforeUpdate()
  updateUpdatedAt(): void {
    this.updatedAt = new Date();
  }

  @BeforeInsert()
  updateCreatedAt(): void {
    this.createdAt = new Date();
  }
}
