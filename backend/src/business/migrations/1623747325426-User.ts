import { MigrationInterface, QueryRunner, Table } from "typeorm";

export class User1623747325426 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: "users",
        columns: [
          {
            name: "id",
            type: "VARCHAR(21)",
            isPrimary: true,
            isNullable: false,
          },
          {
            name: "email",
            type: "VARCHAR(128)",
            isNullable: false,
          },
          {
            name: "username",
            type: "VARCHAR(24)",
            isNullable: false,
          },
          {
            name: "password",
            type: "VARCHAR(72)",
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
    await queryRunner.dropTable("users");
  }
}
