import { MigrationInterface, QueryRunner, Table } from "typeorm";

export class Post1623749528775 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: "posts",
        columns: [
          {
            name: "id",
            type: "VARCHAR(21)",
            isPrimary: true,
            isNullable: false,
          },
          {
            name: "user_id",
            type: "VARCHAR(128)",
            isNullable: false,
          },
          {
            name: "content",
            type: "TEXT",
            isNullable: false,
          },
          {
            name: "created_at",
            type: "TIMESTAMP",
            isNullable: false,
            default: "NOW()",
          },
          {
            name: "updated_at",
            type: "TIMESTAMP",
            isNullable: false,
            default: "NOW()",
          },
        ],
      }),
      true
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropTable("posts");
  }
}
