import os
import json
import psycopg2
from psycopg2 import sql
import boto3

def drop_all_tables(connection):
    try:
        with connection.cursor() as cursor:
            # Retrieve all table names in all schemas except system schemas
            cursor.execute("""
                SELECT table_schema, table_name
                FROM information_schema.tables
                WHERE table_type = 'BASE TABLE'
                  AND table_schema NOT IN ('pg_catalog', 'information_schema');
            """)
            tables = cursor.fetchall()

            if not tables:
                print("No tables found to drop.")
                return

            # Drop each table
            for schema, table in tables:
                drop_query = sql.SQL("DROP TABLE IF EXISTS {schema}.{table} CASCADE;").format(
                    schema=sql.Identifier(schema),
                    table=sql.Identifier(table)
                )
                print(f"Dropping table: {schema}.{table}")
                cursor.execute(drop_query)

            # Commit the changes
            connection.commit()
            print("All tables have been dropped successfully.")

    except Exception as e:
        connection.rollback()
        print(f"An error occurred while dropping tables: {e}")
        raise

def lambda_handler(event, context):
    # Fetch environment variables
    db_host = os.environ['DB_HOST']
    db_port = int(os.environ.get('DB_PORT', 5432))
    db_username = os.environ['DB_USERNAME']
    db_password = os.environ['DB_PASSWORD']
    target_db = os.environ['TARGET_DB']
    ecs_cluster = os.environ['ECS_CLUSTER']
    ecs_service = os.environ['ECS_SERVICE']

    try:
        # Connect to the target database
        conn = psycopg2.connect(
            host=db_host,
            port=db_port,
            user=db_username,
            password=db_password,
            dbname=target_db
        )
        conn.autocommit = False
        print(f"Connected to the database '{target_db}'.")

        # Drop all tables in the target database
        drop_all_tables(conn)

        # Close the database connection
        conn.close()
        print("Database connection closed.")

        # Update ECS service to apply any new changes
        ecs_client = boto3.client('ecs')
        response = ecs_client.update_service(
            cluster=ecs_cluster,
            service=ecs_service,
            forceNewDeployment=True
        )
        print(f"ECS service '{ecs_service}' updated. Response: {response}")

        return {
            'statusCode': 200,
            'body': json.dumps(f"All tables in database '{target_db}' dropped and ECS service '{ecs_service}' restarted successfully.")
        }

    except Exception as e:
        error_message = f"Error processing request: {str(e)}"
        print(error_message)
        return {
            'statusCode': 500,
            'body': json.dumps(error_message)
        }
