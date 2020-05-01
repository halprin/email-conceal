resource "aws_ecs_cluster" "cluster" {
  name = "email-conceal-${var.environment}"

  tags = {
    environment = var.environment
  }
}

resource "aws_ecs_task_definition" "forwarder_task" {
  family = "email-conceal-forwarder-${var.environment}"

  task_role_arn      = aws_iam_role.ecs_task_role.arn
  execution_role_arn = data.aws_iam_role.execution_role.arn

  requires_compatibilities = ["FARGATE"]

  container_definitions = templatefile("${path.module}/forwarderContainerDefinition.json", {
    environment    = var.environment,
    image          = var.docker_image,
    forwarderEmail = "${var.forward_email_prefix}@${var.domain}",
    receivingEmail = var.receiving_email,
    sqsQueueName   = aws_sqs_queue.email_storage_add_event_queue.name,
    region         = data.aws_region.current.name
  })

  network_mode = "awsvpc"
  cpu          = "256"
  memory       = "512"

  tags = {
    environment = var.environment
  }
}

resource "aws_ecs_service" "forwarder_on_cluster" {
  name        = "email-conceal-forwarder-${var.environment}"
  cluster     = aws_ecs_cluster.cluster.id
  launch_type = "FARGATE"

  task_definition = aws_ecs_task_definition.forwarder_task.arn

  desired_count                      = 1
  deployment_minimum_healthy_percent = 100
  deployment_maximum_percent         = 200

  network_configuration {
    subnets          = data.aws_subnet_ids.public_subnets.ids
    security_groups  = [aws_security_group.disallow_all_inbound_traffic.id]
    assign_public_ip = true
  }
}

resource "aws_security_group" "disallow_all_inbound_traffic" {
  name        = "email-conceal-forwarder-${var.environment}"
  description = "Disallow all inbound traffic but allow all outbound"

  vpc_id = data.aws_vpc.main_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    environment = var.environment
  }
}

data "aws_vpc" "main_vpc" {
  default = true
}

data "aws_subnet_ids" "public_subnets" {
  vpc_id = data.aws_vpc.main_vpc.id
  tags = {
    Name = "Public-*"
  }
}
