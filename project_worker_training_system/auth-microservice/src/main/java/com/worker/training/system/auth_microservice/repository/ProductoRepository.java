package com.worker.training.system.auth_microservice.repository;

import com.worker.training.system.auth_microservice.entity.ProductEntity;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface ProductoRepository  extends MongoRepository<ProductEntity,String> {
}
