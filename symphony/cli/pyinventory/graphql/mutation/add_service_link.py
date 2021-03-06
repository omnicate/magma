#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin

from ..fragment.service import ServiceFragment, QUERY as ServiceFragmentQuery

QUERY: List[str] = ServiceFragmentQuery + ["""
mutation AddServiceLinkMutation($id: ID!, $linkId: ID!) {
  addServiceLink(id: $id, linkId: $linkId) {
    ...ServiceFragment
  }
}

"""]

@dataclass
class AddServiceLinkMutation(DataClassJsonMixin):
    @dataclass
    class AddServiceLinkMutationData(DataClassJsonMixin):
        @dataclass
        class Service(ServiceFragment):
            pass

        addServiceLink: Service

    data: AddServiceLinkMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str, linkId: str) -> AddServiceLinkMutationData.Service:
        # fmt: off
        variables = {"id": id, "linkId": linkId}
        try:
            network_start = perf_counter()
            response_text = client.call(''.join(set(QUERY)), variables=variables)
            decode_start = perf_counter()
            res = cls.from_json(response_text).data
            decode_time = perf_counter() - decode_start
            network_time = decode_start - network_start
            client.reporter.log_successful_operation("AddServiceLinkMutation", variables, network_time, decode_time)
            return res.addServiceLink
        except OperationException as e:
            raise FailedOperationException(
                client.reporter,
                e.err_msg,
                e.err_id,
                "AddServiceLinkMutation",
                variables,
            )
