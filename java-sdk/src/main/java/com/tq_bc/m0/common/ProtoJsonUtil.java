package com.tq_bc.m0.common;

import com.google.protobuf.Descriptors;
import com.google.protobuf.Message;
import com.google.protobuf.util.JsonFormat;

import java.io.IOException;
import java.util.List;

public class ProtoJsonUtil {

    private final JsonFormat.Printer printer;
    private final JsonFormat.Parser parser;

    public ProtoJsonUtil() {
        printer = JsonFormat.printer();
        parser = JsonFormat.parser();
    }

    public ProtoJsonUtil(List<Descriptors.Descriptor> anyFieldDescriptor) {
        JsonFormat.TypeRegistry typeRegistry = JsonFormat.TypeRegistry.newBuilder().add(anyFieldDescriptor).build();
        printer = JsonFormat.printer().usingTypeRegistry(typeRegistry);
        parser = JsonFormat.parser().usingTypeRegistry(typeRegistry);
    }

    public String toJson(Message sourceMessage) throws IOException {
        String json = printer.print(sourceMessage);
        return json;
    }

    public Message fromJson(Message.Builder targetBuilder, String json) throws IOException {
        parser.merge(json, targetBuilder);
        return targetBuilder.build();
    }
}
