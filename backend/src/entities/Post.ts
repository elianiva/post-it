import {
  BaseEntity,
  Column,
  Entity,
  PrimaryColumn,
  CreateDateColumn,
  UpdateDateColumn,
} from "typeorm";

@Entity()
export class Post extends BaseEntity {
  @PrimaryColumn({ nullable: false })
  id!: string;

  @Column({ nullable: false })
  user_id!: string;

  @Column({ nullable: false })
  content!: string;

  @CreateDateColumn({ name: "created_at" })
  createdAt!: Date;

  @UpdateDateColumn({ name: "updated_at" })
  updatedAt!: Date;
}
