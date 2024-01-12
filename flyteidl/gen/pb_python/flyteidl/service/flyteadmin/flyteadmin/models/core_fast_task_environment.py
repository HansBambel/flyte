# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class CoreFastTaskEnvironment(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'queue_id': 'str',
        'namespace': 'str',
        'pod_name': 'str'
    }

    attribute_map = {
        'queue_id': 'queue_id',
        'namespace': 'namespace',
        'pod_name': 'pod_name'
    }

    def __init__(self, queue_id=None, namespace=None, pod_name=None):  # noqa: E501
        """CoreFastTaskEnvironment - a model defined in Swagger"""  # noqa: E501

        self._queue_id = None
        self._namespace = None
        self._pod_name = None
        self.discriminator = None

        if queue_id is not None:
            self.queue_id = queue_id
        if namespace is not None:
            self.namespace = namespace
        if pod_name is not None:
            self.pod_name = pod_name

    @property
    def queue_id(self):
        """Gets the queue_id of this CoreFastTaskEnvironment.  # noqa: E501


        :return: The queue_id of this CoreFastTaskEnvironment.  # noqa: E501
        :rtype: str
        """
        return self._queue_id

    @queue_id.setter
    def queue_id(self, queue_id):
        """Sets the queue_id of this CoreFastTaskEnvironment.


        :param queue_id: The queue_id of this CoreFastTaskEnvironment.  # noqa: E501
        :type: str
        """

        self._queue_id = queue_id

    @property
    def namespace(self):
        """Gets the namespace of this CoreFastTaskEnvironment.  # noqa: E501


        :return: The namespace of this CoreFastTaskEnvironment.  # noqa: E501
        :rtype: str
        """
        return self._namespace

    @namespace.setter
    def namespace(self, namespace):
        """Sets the namespace of this CoreFastTaskEnvironment.


        :param namespace: The namespace of this CoreFastTaskEnvironment.  # noqa: E501
        :type: str
        """

        self._namespace = namespace

    @property
    def pod_name(self):
        """Gets the pod_name of this CoreFastTaskEnvironment.  # noqa: E501


        :return: The pod_name of this CoreFastTaskEnvironment.  # noqa: E501
        :rtype: str
        """
        return self._pod_name

    @pod_name.setter
    def pod_name(self, pod_name):
        """Sets the pod_name of this CoreFastTaskEnvironment.


        :param pod_name: The pod_name of this CoreFastTaskEnvironment.  # noqa: E501
        :type: str
        """

        self._pod_name = pod_name

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreFastTaskEnvironment, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreFastTaskEnvironment):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other