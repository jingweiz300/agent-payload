// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: proto/logs/agent_logs_payload.proto

package com.dd.agent.pb;

public interface LogOrBuilder extends
    // @@protoc_insertion_point(interface_extends:pb.Log)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string message = 1;</code>
   * @return The message.
   */
  java.lang.String getMessage();
  /**
   * <code>string message = 1;</code>
   * @return The bytes for message.
   */
  com.google.protobuf.ByteString
      getMessageBytes();

  /**
   * <pre>
   * previously known as "severity"
   * </pre>
   *
   * <code>string status = 2;</code>
   * @return The status.
   */
  java.lang.String getStatus();
  /**
   * <pre>
   * previously known as "severity"
   * </pre>
   *
   * <code>string status = 2;</code>
   * @return The bytes for status.
   */
  com.google.protobuf.ByteString
      getStatusBytes();

  /**
   * <code>int64 timestamp = 3;</code>
   * @return The timestamp.
   */
  long getTimestamp();

  /**
   * <pre>
   * from host
   * </pre>
   *
   * <code>string hostname = 4;</code>
   * @return The hostname.
   */
  java.lang.String getHostname();
  /**
   * <pre>
   * from host
   * </pre>
   *
   * <code>string hostname = 4;</code>
   * @return The bytes for hostname.
   */
  com.google.protobuf.ByteString
      getHostnameBytes();

  /**
   * <pre>
   * from config
   * </pre>
   *
   * <code>string service = 5;</code>
   * @return The service.
   */
  java.lang.String getService();
  /**
   * <pre>
   * from config
   * </pre>
   *
   * <code>string service = 5;</code>
   * @return The bytes for service.
   */
  com.google.protobuf.ByteString
      getServiceBytes();

  /**
   * <code>string source = 6;</code>
   * @return The source.
   */
  java.lang.String getSource();
  /**
   * <code>string source = 6;</code>
   * @return The bytes for source.
   */
  com.google.protobuf.ByteString
      getSourceBytes();

  /**
   * <pre>
   * from config, container tags, ...
   * </pre>
   *
   * <code>repeated string tags = 7;</code>
   * @return A list containing the tags.
   */
  java.util.List<java.lang.String>
      getTagsList();
  /**
   * <pre>
   * from config, container tags, ...
   * </pre>
   *
   * <code>repeated string tags = 7;</code>
   * @return The count of tags.
   */
  int getTagsCount();
  /**
   * <pre>
   * from config, container tags, ...
   * </pre>
   *
   * <code>repeated string tags = 7;</code>
   * @param index The index of the element to return.
   * @return The tags at the given index.
   */
  java.lang.String getTags(int index);
  /**
   * <pre>
   * from config, container tags, ...
   * </pre>
   *
   * <code>repeated string tags = 7;</code>
   * @param index The index of the value to return.
   * @return The bytes of the tags at the given index.
   */
  com.google.protobuf.ByteString
      getTagsBytes(int index);
}
